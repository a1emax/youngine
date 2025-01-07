package input

// KeyboardKeyCode identifies keyboard key.
type KeyboardKeyCode int

const (
	_ KeyboardKeyCode = iota
	KeyboardKeyCodeA
	KeyboardKeyCodeB
	KeyboardKeyCodeC
	KeyboardKeyCodeD
	KeyboardKeyCodeE
	KeyboardKeyCodeF
	KeyboardKeyCodeG
	KeyboardKeyCodeH
	KeyboardKeyCodeI
	KeyboardKeyCodeJ
	KeyboardKeyCodeK
	KeyboardKeyCodeL
	KeyboardKeyCodeM
	KeyboardKeyCodeN
	KeyboardKeyCodeO
	KeyboardKeyCodeP
	KeyboardKeyCodeQ
	KeyboardKeyCodeR
	KeyboardKeyCodeS
	KeyboardKeyCodeT
	KeyboardKeyCodeU
	KeyboardKeyCodeV
	KeyboardKeyCodeW
	KeyboardKeyCodeX
	KeyboardKeyCodeY
	KeyboardKeyCodeZ
	KeyboardKeyCodeAltLeft
	KeyboardKeyCodeAltRight
	KeyboardKeyCodeArrowDown
	KeyboardKeyCodeArrowLeft
	KeyboardKeyCodeArrowRight
	KeyboardKeyCodeArrowUp
	KeyboardKeyCodeBackquote
	KeyboardKeyCodeBackslash
	KeyboardKeyCodeBackspace
	KeyboardKeyCodeBracketLeft
	KeyboardKeyCodeBracketRight
	KeyboardKeyCodeCapsLock
	KeyboardKeyCodeComma
	KeyboardKeyCodeContextMenu
	KeyboardKeyCodeControlLeft
	KeyboardKeyCodeControlRight
	KeyboardKeyCodeDelete
	KeyboardKeyCodeDigit0
	KeyboardKeyCodeDigit1
	KeyboardKeyCodeDigit2
	KeyboardKeyCodeDigit3
	KeyboardKeyCodeDigit4
	KeyboardKeyCodeDigit5
	KeyboardKeyCodeDigit6
	KeyboardKeyCodeDigit7
	KeyboardKeyCodeDigit8
	KeyboardKeyCodeDigit9
	KeyboardKeyCodeEnd
	KeyboardKeyCodeEnter
	KeyboardKeyCodeEqual
	KeyboardKeyCodeEscape
	KeyboardKeyCodeF1
	KeyboardKeyCodeF2
	KeyboardKeyCodeF3
	KeyboardKeyCodeF4
	KeyboardKeyCodeF5
	KeyboardKeyCodeF6
	KeyboardKeyCodeF7
	KeyboardKeyCodeF8
	KeyboardKeyCodeF9
	KeyboardKeyCodeF10
	KeyboardKeyCodeF11
	KeyboardKeyCodeF12
	KeyboardKeyCodeHome
	KeyboardKeyCodeInsert
	KeyboardKeyCodeMetaLeft
	KeyboardKeyCodeMetaRight
	KeyboardKeyCodeMinus
	KeyboardKeyCodeNumLock
	KeyboardKeyCodeNumpad0
	KeyboardKeyCodeNumpad1
	KeyboardKeyCodeNumpad2
	KeyboardKeyCodeNumpad3
	KeyboardKeyCodeNumpad4
	KeyboardKeyCodeNumpad5
	KeyboardKeyCodeNumpad6
	KeyboardKeyCodeNumpad7
	KeyboardKeyCodeNumpad8
	KeyboardKeyCodeNumpad9
	KeyboardKeyCodeNumpadAdd
	KeyboardKeyCodeNumpadDecimal
	KeyboardKeyCodeNumpadDivide
	KeyboardKeyCodeNumpadEnter
	KeyboardKeyCodeNumpadEqual
	KeyboardKeyCodeNumpadMultiply
	KeyboardKeyCodeNumpadSubtract
	KeyboardKeyCodePageDown
	KeyboardKeyCodePageUp
	KeyboardKeyCodePause
	KeyboardKeyCodePeriod
	KeyboardKeyCodePrintScreen
	KeyboardKeyCodeQuote
	KeyboardKeyCodeScrollLock
	KeyboardKeyCodeSemicolon
	KeyboardKeyCodeShiftLeft
	KeyboardKeyCodeShiftRight
	KeyboardKeyCodeSlash
	KeyboardKeyCodeSpace
	KeyboardKeyCodeTab

	// MaxKeyboardKeyCode specifies the maximum known value of the [KeyboardKeyCode] type.
	MaxKeyboardKeyCode KeyboardKeyCode = iota - 1

	// MinKeyboardKeyCode specifies the minimum known value of the [KeyboardKeyCode] type.
	MinKeyboardKeyCode KeyboardKeyCode = 1

	// KeyboardKeyCount specifies the number of known keyboard keys.
	KeyboardKeyCount = int(MaxKeyboardKeyCode)
)
