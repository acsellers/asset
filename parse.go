package asset

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type require struct {
	Type, Item string
}

func parseSheet(content string) ([]*require, []string) {
	scanner := bufio.NewScanner(bytes.NewBufferString(content))
	currentStyle, requires := []string{}, []*require{}
	var inComment bool
	var trimText string
	for scanner.Scan() {
		trimText = strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(trimText, "//=") {
			comm, _ := getCommand(trimText)
			if comm != nil {
				requires = append(requires, comm)
			} else {
				currentStyle = append(currentStyle, trimText)
			}
		} else {

			if strings.Contains(trimText, "/*") {

				if strings.Contains(trimText[strings.Index(trimText, "/*"):], "*/") {
					currentStyle = append(currentStyle, trimText)

				} else {
					inComment = true

					for inComment {
						if strings.HasPrefix(trimText, "*=") {
							comm, _ := getCommand(trimText)
							if comm != nil {
								requires = append(requires, comm)
							} else {
								currentStyle = append(currentStyle, trimText)
							}
						} else {
							if trimText != "" {
								currentStyle = append(currentStyle, trimText)
							}
						}

						inComment = scanner.Scan()
						if inComment {
							trimText = strings.TrimSpace(scanner.Text())
							if trimText == "*/" {
								currentStyle = append(currentStyle, "*/")
								inComment = false
							}
						}
					}
				}
			} else {
				if trimText != "" {
					currentStyle = append(currentStyle, trimText)
				}
			}
		}
	}
	return requires, currentStyle
}

func getCommand(c string) (*require, error) {
	c = extractPrefix(c)
	parts := strings.SplitN(c, " ", 2)
	switch parts[0] {
	case "require":
		if len(parts) == 2 {
			return &require{Item: strings.Trim(parts[1], `"'`)}, nil
		}
	case "require_self":
		return &require{Type: "self"}, nil
	case "require_tree":
		if len(parts) == 2 {
			return &require{Item: strings.Trim(parts[1], `"'`), Type: "tree"}, nil
		}
	case "require_directory":
		if len(parts) == 2 {
			return &require{Item: strings.Trim(parts[1], `"'`), Type: "directory"}, nil
		}
	}
	return nil, fmt.Errorf("Could not recognize command")
}

func extractPrefix(c string) string {
	return strings.Trim(c, "/=* ")
}
