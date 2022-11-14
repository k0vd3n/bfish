# BlowFish
The program of encryption and decryption according to the BlowFish algorithm

# Description
bfish is an application for symmetric encryption and decryption of text, hex, decimal data. Blowfish is a cryptographic algorithm that implements block symmetric encryption with a variable key length. 
	
Designed by Bruce Schneierin 1993. It is a Feistel network. Performed on simple and fast operations: XOR, substitution, addition. It is non-proprietary and freely distributed. You can readmore about how the algorithm works on the Internet

The program can show the current value of the main key, the round keys of the boxes and the number of rounds
 
The program can change the values of the boxes, round keys, number of rounds and the main key 

# Encryption
Encryption is available in two modes: in 64-bit mode and in full-string encryption mode.

In the 64-bit block encryption mode, 3 data input options are available: decimal, hexadecimal, and string. The result of encryption of this mode will be an array of bytes, 2 numbers in hexadecimal and decimal form and as a string.

In full-string encryption mode (full) data entry is only available as a string of unlimited length, as the name implies. The result of this encryption will be an array of bytes

# Decryption
Decryption is available in two modes: in 64-bit mode and in full-string decryption mode.

In the decryption mode of a 64-bit block, 3 data input options are available: decimal, hexadecimal, and string. The result of encryption of this mode will be an array of bytes, 2 numbers in hexadecimal and decimal form and as a string.

In the decryption mode of the full string (full) data input is available only in a byte array of unlimited length. The result of this encryption will be a string and an array of bytes

# Examples of commands
To encrypt a block of 64 bits, you can run the following command, in which the characters in quotation marks are a message divided into 32 bits

`bfish encrypts the string "exam" "ple1" -m`

It should be borne in mind that not all characters weigh 1 byte.
You can also use the hexadecimal input mode 

`bfish encrypts "1234" "5678" -m hexadecimal code`

The input will be a decimal number, but the value of the flag must be equal to decimal

In full-string encryption mode, to encrypt a message, you need to run the following command, where the characters in quotation marks are the message for encryption

`bfish encrypts the full "Hello, world!"`

To decrypt a block of 64 bits, you can run the following command, in which the characters in quotation marks are a message divided into 32 bits

`bfish decrypt "51eaf5b2" "ba6b19dd" -m hexadecimal`

You can also use the decimal number input mode 

`bfish decrypts "901818609" "3006639247" -m decimal number`

The string will be the same as the input string, but the value of the flag must be equal. It should be borne in mind that not all characters weigh 1 byte.

The program can output the value of the parameter involved in encryption, for example:

`bfish gets the key`

However, the program has not been finalized yet and in order for the parameter value to be output, it is necessary to generate it in the directory in which you are currently located

It is possible to change the parameters involved in encryption and decryption:

`bfish change the "blowfish" key`
`bfish change mailbox 111111 0 1` 

When changing the sbox information or the pkey key, the first parameter is the new value, the second is the sbox information number or index of the pkey key element, the third value is used only when changing the sbox information, it is equal to the sbox information of the element that you want to change 

More detailed instructions for each command are available in the program
