package types

const RUNNING = true
const PAUSED = false
const PROGRAM_ADDR = 0x200

type Chip8 struct {
	Emu_state   bool
	V           [16]byte
	Ram         [4096]byte
	Display     [64 * 32]bool
	PC          uint16
	I           uint16
	Stack       [16]uint16
	Sp          int16
	Sound_Timer byte
	Delay_Timer byte
}
