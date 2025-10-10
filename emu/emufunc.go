package emu

import (
	"bufio"
	"go8/opcodes"
	"go8/types"
	"log"
	"math/rand"
	"os"
)

var was_key_pressed bool = false
var previous_key_index int = 0

func Load_rom(filename string, chip8 *types.Chip8) {
	rom_file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(rom_file)
	buffer := make([]byte, 4026)
	_, err = reader.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	for i := types.PROGRAM_ADDR; i < len(buffer); i++ {
		chip8.Ram[i] = buffer[i-types.PROGRAM_ADDR]
	}
	index := types.FONT_ADDR
	for _, font_byte := range types.Font_Set {
		chip8.Ram[index] = font_byte
		index++
	}
	chip8.Emu_state = types.RUNNING
	buffer = nil
	chip8.PC = types.PROGRAM_ADDR
}

func Chip8_cycle(chip8 *types.Chip8) {
	if chip8.Delay_Timer > 0 {
		chip8.Delay_Timer--
	}
	if chip8.Sound_Timer > 0 {
		chip8.Sound_Timer--
	}
	op := fetch_op(chip8)
	decode_op(op, chip8)
}

func fetch_op(chip *types.Chip8) (op uint16) {
	var first_byte byte = chip.Ram[chip.PC]
	var second_byte byte = chip.Ram[chip.PC+1]
	chip.PC += 2
	temp_1 := uint16(first_byte)
	temp_2 := uint16(second_byte)
	op = temp_1<<8 | temp_2
	return
}

func decode_op(op uint16, chip *types.Chip8) {
	if op == 0x00E0 {
		for y := range 32 {
			for x := range 64 {
				chip.Display[y][x] = false
			}
		}
		return
	}
	if op == 0x00EE {
		chip.Sp--
		chip.PC = chip.Stack[chip.Sp]
		return
	}
	op_type := (op & 0xF000)
	switch op_type {
	case opcodes.DRAW:
		sprite_address := chip.I
		length := (op & 0x000F)
		x_reg := (op & 0x0F00) >> 8
		y_reg := (op & 0x00F0) >> 4
		x_pos := chip.V[x_reg] % 64
		y_pos := chip.V[y_reg] % 32
		y_len := 0
		for y := y_pos; byte(y_len) < byte(length); y++ {
			x_len := 0
			if y == 32 {
				return
			}
			for x := x_pos; x_len < 8; x++ {
				sprite_byte := chip.Ram[sprite_address]
				sprite_bit := sprite_byte & (0x80 >> x_len)
				if x >= 64 {
					break
				}
				if sprite_bit != 0 {
					if chip.Display[y][x] == true {
						chip.Display[y][x] = false
						chip.V[15] = 1
					} else {
						chip.Display[y][x] = true
					}
				}
				x_len++
			}
			sprite_address++
			y_len++

		}
	case opcodes.JUMP:
		address := (op & 0x0FFF)
		chip.PC = address
	case opcodes.JUMP_WITH_OFF:
		addr := op & 0x0FFF
		chip.PC = addr + uint16(chip.V[0])
	case opcodes.CALL:
		sub_address := op & 0x0FFF
		chip.Stack[chip.Sp] = chip.PC
		chip.Sp++
		chip.PC = sub_address
	case opcodes.ADD:
		reg_addres := (op & 0x0F00) >> 8
		add_value := (op & 0x00FF)
		chip.V[reg_addres] += byte(add_value)
	case opcodes.SET_I:
		adress_shift := op << 4
		chip.I = adress_shift >> 4
	case opcodes.SET:
		reg_addres := (op & 0x0F00) >> 8
		reg_value := (op & 0x00FF)
		chip.V[reg_addres] = byte(reg_value)
	case opcodes.SKIP_IF:
		reg_addres := (op & 0x0F00) >> 8
		comp_value := (op & 0x0FF)
		if chip.V[reg_addres] == byte(comp_value) {
			chip.PC += 2
		}
	case opcodes.SKIP_NOT:
		reg_addres := (op & 0x0F00) >> 8
		comp_value := (op & 0x0FF)
		if chip.V[reg_addres] != byte(comp_value) {
			chip.PC += 2
		}
	case opcodes.SKIP_REG:
		reg_x := (op & 0x0F00) >> 8
		reg_y := (op & 0x00F0) >> 4
		if chip.V[reg_x] == chip.V[reg_y] {
			chip.PC += 2
		}
	case opcodes.SKIP_NOT_REG:
		reg_x := (op & 0x0F00) >> 8
		reg_y := (op & 0x00F0) >> 4
		if chip.V[reg_x] != chip.V[reg_y] {
			chip.PC += 2
		}
	case opcodes.RAND:
		reg_addr := (op & 0x0F00) >> 8
		bit_and_value := (op & 0x00FF)
		rand_num := byte(rand.Uint32() % 255)
		chip.V[reg_addr] = (rand_num & byte(bit_and_value))
	case opcodes.REG_INSTRUCT:
		handle_reg_instruct(op, chip)
	case opcodes.F_INSTRUCT:
		handle_F_instructs(op, chip)
	case opcodes.E_INSTRUCT:
		bottom_nibble := op & 0x00FF
		switch bottom_nibble {
		case 0x009E:
			reg_addr := (op & 0x0F00) >> 8
			reg_value := chip.V[reg_addr]
			if chip.Key_board[reg_value] {
				chip.PC += 2
			}

		case 0x00A1:
			reg_addr := (op & 0x0F00) >> 8
			reg_value := chip.V[reg_addr]
			if chip.Key_board[reg_value] == false {
				chip.PC += 2
			}
		}

	}

}
func handle_reg_instruct(op uint16, chip *types.Chip8) {
	op_type := op & 0x000F
	switch op_type {
	case opcodes.REG_SET:
		reg_addr := (op & 0x0F00) >> 8
		reg_y_addr := (op & 0x00F0) >> 4
		chip.V[reg_addr] = chip.V[reg_y_addr]
	case opcodes.REG_BIN_OR:
		reg_addr := (op & 0x0F00) >> 8
		reg_y_addr := (op & 0x00F0) >> 4
		chip.V[reg_addr] = (chip.V[reg_addr] | chip.V[reg_y_addr])
		chip.V[0xF] = 0
	case opcodes.REG_BIN_AND:
		reg_addr := (op & 0x0F00) >> 8
		reg_y_addr := (op & 0x00F0) >> 4
		chip.V[reg_addr] = (chip.V[reg_addr] & chip.V[reg_y_addr])
		chip.V[0xF] = 0

	case opcodes.REG_XOR:
		reg_y_addr := (op & 0x00F0) >> 4
		reg_addr := (op & 0x0F00) >> 8
		chip.V[reg_addr] = (chip.V[reg_addr] ^ chip.V[reg_y_addr])
		chip.V[0xF] = 0
	case opcodes.REG_ADD:
		reg_addr := (op & 0x0F00) >> 8
		reg_y_addr := (op & 0x00F0) >> 4
		x_value := chip.V[reg_addr]
		y_value := chip.V[reg_y_addr]
		chip.V[reg_addr] = (y_value + x_value)
		if chip.V[reg_addr] < x_value {
			chip.V[0xF] = 1
		} else {
			chip.V[0xF] = 0
		}
	case opcodes.REG_SUB_X_Y:
		reg_addr := (op & 0x0F00) >> 8
		reg_addr_y := (op & 0x00F0) >> 4
		x_value := chip.V[reg_addr]
		y_value := chip.V[reg_addr_y]
		chip.V[reg_addr] = (x_value - y_value)
		if x_value < y_value {
			chip.V[0xF] = 0
		} else {
			chip.V[0xF] = 1
		}

	case opcodes.REG_SUB_Y_X:
		reg_addr := (op & 0x0F00) >> 8
		reg_addr_y := (op & 0x00F0) >> 4
		x_value := chip.V[reg_addr]
		y_value := chip.V[reg_addr_y]
		chip.V[reg_addr] = (y_value - x_value)
		if y_value < x_value {
			chip.V[0xF] = 0
		} else {
			chip.V[0xF] = 1
		}

	case opcodes.REG_SHIFT_R:
		reg_addr := (op & 0x0F00) >> 8
		reg_addr_y := (op & 0x00F0) >> 4
		//make this action configurable my user at some point
		bit_to_be_shifted := chip.V[reg_addr_y] & 0x01
		chip.V[reg_addr] = (chip.V[reg_addr_y] >> 1)
		chip.V[0xF] = bit_to_be_shifted
	case opcodes.REG_SHIFT_L:
		reg_addr := (op & 0x0F00) >> 8
		reg_addr_y := (op & 0x00F0) >> 4
		//make this action configurable my user at some point
		//	chip.V[reg_addr] = chip.V[reg_addr_y]
		bit_to_be_shifted := (chip.V[reg_addr_y] & 0x80) >> 7
		chip.V[reg_addr] = (chip.V[reg_addr_y] << 1)
		chip.V[0xF] = bit_to_be_shifted

	}
}
func handle_F_instructs(op uint16, chip *types.Chip8) {
	op_type := op & 0x00FF
	reg_addr := (op & 0x0F00) >> 8
	switch op_type {
	case opcodes.STORE_REG_TO_MEM:
		for i := 0; i <= int(reg_addr); i++ {
			chip.Ram[chip.I] = chip.V[i]
			chip.I++
		}
	case opcodes.LOAD_FROM_MEM:
		for i := 0; i <= int(reg_addr); i++ {
			chip.V[i] = chip.Ram[chip.I]
			chip.I++
		}
	case opcodes.ADD_TO_I:
		index_value := uint16(chip.V[reg_addr])
		chip.I += index_value
	case opcodes.BIN_CODED:

		value := chip.V[reg_addr]
		chip.Ram[chip.I+2] = value % 10
		value /= 10
		chip.Ram[chip.I+1] = value % 10
		value /= 10
		chip.Ram[chip.I] = value % 10
	case opcodes.GET_KEY:
		reg_addr := (op & 0x0F00) >> 8
		if was_key_pressed {
			if !chip.Key_board[previous_key_index] {
				was_key_pressed = false
				chip.V[reg_addr] = byte(previous_key_index)
				previous_key_index = 0
				return
			}
		}
		for key_index, key_down := range chip.Key_board {
			if key_down {
				was_key_pressed = true
				previous_key_index = key_index

			}
		}
		chip.PC -= 2
	case opcodes.FONT_GET:
		x_value := chip.V[reg_addr]
		chip.I = uint16(x_value)
	case opcodes.SET_VX_TO_DELAY:
		chip.V[reg_addr] = chip.Delay_Timer

	case opcodes.SET_DELAY:
		chip.Delay_Timer = chip.V[reg_addr]
	case opcodes.SET_SOUND:
		chip.Sound_Timer = chip.V[reg_addr]
	}
}
