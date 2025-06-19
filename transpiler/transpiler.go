package transpiler

import (
	"regexp"
	"strings"
)

// trimTrailingBrace entfernt ein `{` am Ende einer Zeile, falls vorhanden
func trimTrailingBrace(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasSuffix(s, "{") {
		return strings.TrimSpace(s[:len(s)-1])
	}
	return s
}

// functionHasReturn prüft, ob im nächsten Block ein return vorkommt
// und bestimmt den Rückgabetyp als String (z.B. "void", "int" oder "String")
func functionHasReturn(lines []string, start int) (bool, string) {
	depth := 0
	for i := start; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if strings.HasSuffix(line, "{") {
			depth++
		}
		if strings.HasPrefix(line, "return ") {
			returnExpr := strings.TrimSpace(line[len("return "):])
			// Typ ableiten: wenn return Zahl (nur Ziffern), dann int, wenn String in "", dann String, sonst void
			if strings.HasPrefix(returnExpr, "\"") && strings.HasSuffix(returnExpr, "\"") {
				return true, "String"
			}
			// check ob Zahl (int)
			isInt := true
			for _, ch := range returnExpr {
				if ch < '0' || ch > '9' {
					isInt = false
					break
				}
			}
			if isInt {
				return true, "int"
			}
			// Standard fallback
			return true, "void"
		}
		if line == "}" {
			if depth == 0 {
				break
			}
			depth--
		}
	}
	return false, "void"
}

// Transpile konvertiert LiteJava-Code in gültigen Java-Code
func Transpile(input string) string {
	lines := strings.Split(input, "\n")
	var output []string
	indent := ""

	className := ""
	inMainClass := false // Flag, ob wir in der main-Klasse sind

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		switch {
		case strings.HasPrefix(line, "class "):
			className = trimTrailingBrace(strings.TrimPrefix(line, "class "))
			output = append(output, "public class "+className+" {")
			indent = "  "
			inMainClass = true // angenommen Hauptklasse mit main-Methode

		case strings.HasPrefix(line, "main"):
			output = append(output, indent+"public static void main(String[] args) {")
			indent += "  "

		case strings.HasPrefix(line, "print "):
			text := strings.TrimPrefix(line, "print ")
			output = append(output, indent+"System.out.println("+text+");")

		case strings.HasPrefix(line, "var "):
			// Ersetze var durch int
			re := regexp.MustCompile(`var (\w+) = (.+)`)
			match := re.FindStringSubmatch(line)
			if len(match) == 3 {
				output = append(output, indent+"int "+match[1]+" = "+match[2]+";")
			}

		case strings.HasPrefix(line, "if ("):
			cond := trimTrailingBrace(line)
			if len(cond) >= 4 {
				output = append(output, indent+"if "+cond[3:]+" {")
			} else {
				output = append(output, indent+line+" {")
			}
			indent += "  "

		case strings.HasPrefix(line, "for ("):
			// Ersetze var in for-Schleife durch int
			line = strings.Replace(line, "var ", "int ", 1)
			output = append(output, indent+line)
			if !strings.HasSuffix(line, "{") {
				output[len(output)-1] += " {"
			}
			indent += "  "

		case strings.HasPrefix(line, "func "):
			re := regexp.MustCompile(`func (\w+)\((.*?)\)`)
			match := re.FindStringSubmatch(line)
			if len(match) == 3 {
				funcName := match[1]
				paramsRaw := strings.TrimSpace(match[2])
				params := ""
				if paramsRaw != "" {
					paramList := strings.Split(paramsRaw, ",")
					var typedParams []string
					for _, p := range paramList {
						p = strings.TrimSpace(p)
						if p != "" {
							typedParams = append(typedParams, "String "+p)
						}
					}
					params = strings.Join(typedParams, ", ")
				}

				_, retType := functionHasReturn(lines, i+1)
				if retType == "" {
					retType = "void"
				}

				// Methoden in Hauptklasse als static deklarieren
				staticStr := ""
				if inMainClass {
					staticStr = "static "
				}

				output = append(output, indent+"public "+staticStr+retType+" "+funcName+"("+params+") {")
				indent += "  "
			}

		case strings.HasPrefix(line, "return "):
			// return; - Semikolon ergänzen
			if !strings.HasSuffix(line, ";") {
				line += ";"
			}
			output = append(output, indent+line)

		case line == "}":
			if len(indent) >= 2 {
				indent = indent[:len(indent)-2]
			}
			output = append(output, indent+"}")

		case line == "":
			// Leere Zeile ignorieren

		default:
			trimmed := strings.TrimSpace(line)
			// Zeilen ohne ; oder { oder } am Ende bekommen Semikolon
			if !strings.HasSuffix(trimmed, ";") && !strings.HasSuffix(trimmed, "{") && !strings.HasSuffix(trimmed, "}") {
				output = append(output, indent+line+";")
			} else {
				output = append(output, indent+line)
			}
		}
	}

	// Offene Blöcke am Ende schließen
	for len(indent) > 0 {
		if len(indent) >= 2 {
			indent = indent[:len(indent)-2]
		} else {
			indent = ""
		}
		output = append(output, indent+"}")
	}

	return strings.Join(output, "\n")
}