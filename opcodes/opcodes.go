package opcodes

const CLEAR uint16 = 0x00E0
const JUMP uint16 = 0x1000
const SET uint16 = 0x6000
const ADD uint16 = 0x7000
const RETURN uint16 = 0x0EE
const CALL uint16 = 0x2000
const SET_I uint16 = 0xA000
const DRAW uint16 = 0xD000
const SKIP_IF uint16 = 0x3000
const SKIP_NOT uint16 = 0x4000
const SKIP_REG uint16 = 0x5000
const SKIP_NOT_REG uint16 = 0x9000
const REG_INSTRUCT uint16 = 0x8000
const RAND uint16 = 0xC000

// this is the start of the 0x8000 insturctions these are all logical and math insturctions that only operate on registers

const REG_SET uint16 = 0
const REG_BIN_OR uint16 = 0x0001
const REG_BIN_AND uint16 = 0x0002
const REG_XOR uint16 = 0x0003
const REG_ADD uint16 = 0x0004
const REG_SUB_X_Y uint16 = 0x0005
const REG_SUB_Y_X uint16 = 0x0007
const REG_SHIFT_R uint16 = 0x0006
const REG_SHIFT_L uint16 = 0x000E

// this is the staart of the 0xF000 insturctions this all have to be parsed diffrently than the 0x8000 and the first set of insturctions
const STORE_REG_TO_MEM uint16 = 0x0055
const LOAD_FROM_MEM uint16 = 0x0065
