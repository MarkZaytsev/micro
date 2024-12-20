package action

var bufdefaults = map[string]string{
	"Up":             "CursorUp",
	"Down":           "CursorDown",
	"Right":          "CursorRight",
	"Left":           "CursorLeft",
	"ShiftUp":        "SelectUp",
	"ShiftDown":      "SelectDown",
	"ShiftLeft":      "SelectLeft",
	"ShiftRight":     "SelectRight",
	"AltLeft":        "WordLeft",
	"AltRight":       "WordRight",
	"AltUp":          "MoveLinesUp",
	"AltDown":        "MoveLinesDown",
	"AltShiftRight":  "SelectWordRight",
	"AltShiftLeft":   "SelectWordLeft",
	"CtrlLeft":       "StartOfTextToggle",
	"CtrlRight":      "EndOfLine",
	"CtrlShiftLeft":  "SelectToStartOfTextToggle",
	"ShiftHome":      "SelectToStartOfTextToggle",
	"CtrlShiftRight": "SelectToEndOfLine",
	"ShiftEnd":       "SelectToEndOfLine",
	"CtrlUp":         "CursorStart",
	"CtrlDown":       "CursorEnd",
	"CtrlShiftUp":    "SelectToStart",
	"CtrlShiftDown":  "SelectToEnd",
	"Alt-{":          "ParagraphPrevious",
	"Alt-}":          "ParagraphNext",
	"Enter":          "InsertNewline",
	"CtrlH":          "Backspace",
	"Backspace":      "Backspace",
	"OldBackspace":   "Backspace",
	"Alt-CtrlH":      "DeleteWordLeft",
	"Alt-Backspace":  "DeleteWordLeft",
	"Tab":            "Autocomplete|IndentSelection|InsertTab",
	"Backtab":        "CycleAutocompleteBack|OutdentSelection|OutdentLine",
	"Ctrl-o":         "OpenFile",
	"Ctrl-s":         "Save",
	"Ctrl-f":         "Find",
	"Alt-F":          "FindLiteral",
	"Ctrl-n":         "FindNext",
	"Ctrl-p":         "FindPrevious",
	"Alt-[":          "DiffPrevious|CursorStart",
	"Alt-]":          "DiffNext|CursorEnd",
	"Ctrl-z":         "Undo",
	"Ctrl-y":         "Redo",
	"Ctrl-c":         "Copy|CopyLine",
	"Ctrl-x":         "Cut|CutLine",
	"Ctrl-k":         "CutLineAppend",
	"Ctrl-d":         "Duplicate|DuplicateLine",
	"Ctrl-v":         "Paste",
	"Ctrl-a":         "SelectAll",
	"Ctrl-t":         "AddTab",
	"Alt-,":          "PreviousTab|LastTab",
	"Alt-.":          "NextTab|FirstTab",
	"Home":           "StartOfTextToggle",
	"End":            "EndOfLine",
	"CtrlHome":       "CursorStart",
	"CtrlEnd":        "CursorEnd",
	"PageUp":         "CursorPageUp",
	"PageDown":       "CursorPageDown",
	"CtrlPageUp":     "PreviousTab|LastTab",
	"CtrlPageDown":   "NextTab|FirstTab",
	"ShiftPageUp":    "SelectPageUp",
	"ShiftPageDown":  "SelectPageDown",
	"Ctrl-g":         "ToggleHelp",
	"Alt-g":          "ToggleKeyMenu",
	"Ctrl-r":         "ToggleRuler",
	"Ctrl-l":         "command-edit:goto ",
	"Delete":         "Delete",
	"Ctrl-b":         "ShellMode",
	"Ctrl-q":         "Quit",
	"Ctrl-e":         "CommandMode",
	"Ctrl-w":         "NextSplit|FirstSplit",
	"Ctrl-u":         "ToggleMacro",
	"Ctrl-j":         "PlayMacro",
	"Insert":         "ToggleOverwriteMode",

	// Emacs-style keybindings
	"Alt-f": "WordRight",
	"Alt-b": "WordLeft",
	"Alt-a": "StartOfText",
	"Alt-e": "EndOfLine",
	// "Alt-p": "CursorUp",
	// "Alt-n": "CursorDown",

	// Integration with file managers
	"F2":  "Save",
	"F3":  "Find",
	"F4":  "Quit",
	"F7":  "Find",
	"F10": "Quit",
	"Esc": "Escape,Deselect,ClearInfo,RemoveAllMultiCursors,UnhighlightSearch",

	// Mouse bindings
	"MouseWheelUp":     "ScrollUp",
	"MouseWheelDown":   "ScrollDown",
	"MouseLeft":        "MousePress",
	"MouseLeftDrag":    "MouseDrag",
	"MouseLeftRelease": "MouseRelease",
	"MouseMiddle":      "PastePrimary",
	"Ctrl-MouseLeft":   "MouseMultiCursor",

	"Alt-n":        "SpawnMultiCursor",
	"AltShiftUp":   "SpawnMultiCursorUp",
	"AltShiftDown": "SpawnMultiCursorDown",
	"Alt-m":        "SpawnMultiCursorSelect",
	"Alt-p":        "RemoveMultiCursor",
	"Alt-c":        "RemoveAllMultiCursors",
	"Alt-x":        "SkipMultiCursor",
}

var infodefaults = map[string]string{
	"Up":             "HistoryUp",
	"Down":           "HistoryDown",
	"Right":          "CursorRight",
	"Left":           "CursorLeft",
	"ShiftUp":        "SelectUp",
	"ShiftDown":      "SelectDown",
	"ShiftLeft":      "SelectLeft",
	"ShiftRight":     "SelectRight",
	"AltLeft":        "WordLeft",
	"AltRight":       "WordRight",
	"AltUp":          "CursorStart",
	"AltDown":        "CursorEnd",
	"AltShiftRight":  "SelectWordRight",
	"AltShiftLeft":   "SelectWordLeft",
	"CtrlLeft":       "StartOfTextToggle",
	"CtrlRight":      "EndOfLine",
	"CtrlShiftLeft":  "SelectToStartOfTextToggle",
	"ShiftHome":      "SelectToStartOfTextToggle",
	"CtrlShiftRight": "SelectToEndOfLine",
	"ShiftEnd":       "SelectToEndOfLine",
	"CtrlUp":         "CursorStart",
	"CtrlDown":       "CursorEnd",
	"CtrlShiftUp":    "SelectToStart",
	"CtrlShiftDown":  "SelectToEnd",
	"Enter":          "ExecuteCommand",
	"CtrlH":          "Backspace",
	"Backspace":      "Backspace",
	"OldBackspace":   "Backspace",
	"Alt-CtrlH":      "DeleteWordLeft",
	"Alt-Backspace":  "DeleteWordLeft",
	"Tab":            "CommandComplete",
	"Backtab":        "CycleAutocompleteBack",
	"Ctrl-z":         "Undo",
	"Ctrl-y":         "Redo",
	"Ctrl-c":         "Copy|CopyLine",
	"Ctrl-x":         "Cut|CutLine",
	"Ctrl-k":         "CutLineAppend",
	"Ctrl-v":         "Paste",
	"Home":           "StartOfTextToggle",
	"End":            "EndOfLine",
	"CtrlHome":       "CursorStart",
	"CtrlEnd":        "CursorEnd",
	"Delete":         "Delete",
	"Ctrl-q":         "AbortCommand",
	"Ctrl-e":         "EndOfLine",
	"Ctrl-a":         "StartOfLine",
	"Ctrl-w":         "DeleteWordLeft",
	"Insert":         "ToggleOverwriteMode",
	"Ctrl-b":         "WordLeft",
	"Ctrl-f":         "WordRight",
	"Ctrl-d":         "DeleteWordLeft",
	"Ctrl-m":         "ExecuteCommand",
	"Ctrl-n":         "HistoryDown",
	"Ctrl-p":         "HistoryUp",
	"Ctrl-u":         "SelectToStart",

	// Emacs-style keybindings
	"Alt-f": "WordRight",
	"Alt-b": "WordLeft",
	"Alt-a": "StartOfText",
	"Alt-e": "EndOfLine",

	// Integration with file managers
	"F10": "AbortCommand",
	"Esc": "AbortCommand",

	// Mouse bindings
	"MouseWheelUp":     "HistoryUp",
	"MouseWheelDown":   "HistoryDown",
	"MouseLeft":        "MousePress",
	"MouseLeftDrag":    "MouseDrag",
	"MouseLeftRelease": "MouseRelease",
	"MouseMiddle":      "PastePrimary",
}
