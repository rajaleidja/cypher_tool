# cypher_tool

***


## 1 Week TeamWork.  Cypher_tool

### Nikita Strekalov - Team leader
### Enam Hossain - Member
### Oleh Ilchuk - Member
---

___
> Encryption is the key to understanding, but only if you lost that key in the beginning.
>
> Nikita Strekalov
___

***

### Description

Cypher Tool is a simple command-line tool, 
which is designed to encrypt and decrypt messages using various encryption methods and understanding encryption algorithms.


***


## ROT13


ROT13 is a special case of the Caesar cipher, where the shift equals 13 positions in the alphabet. Caesar's cipher generally assumes a shift of any given number of positions. ROT13 is just a particular variant of this cipher where the shift is always 13.




### What is ROT13


ROT13 is a simple encryption method that relies on shifting each letter in the alphabet by 13 positions. ROT stands for "rotate" and the number 13 is the number of positions by which the shift occurs.


An example of how ROT13 works:


Each letter of the alphabet is assigned a numeric value. For example, 'A' is 1, 'B' is 2, ..., ..., 'Z' is 26.


ROT13 shifts each letter forward by 13 positions. For example, 'A' becomes 'N', 'B' becomes 'O', and so on.


Similarly, if we apply ROT13 to the encrypted text, we get the original text, since shifting 13 positions forward and backward produces the original result.


***
Example:


The original message: "HELLO."


ROT13 encryption: "URYYB".


Decryption ROT13: "HELLO".
***


![ROT13 "HELLO"](https://www.google.com/url?sa=i&url=https%3A%2F%2Fru.wikipedia.org%2Fwiki%2FROT13&psig=AOvVaw1DT03FmJ3kD-k54uB5Gg-j&ust=1699144776105000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCNCDgNiNqYIDFQAAAAAdAAAAABAE)


**ROT13 is a simple and easily reversible encryption method** that is sometimes used to hide text in publicly accessible places, but ***is not intended to seriously protect sensitive information***.


---


### How ROT13 works


The way ROT13 works is that each letter is replaced by a letter that is 13 positions forward. If it reaches the end of the alphabet, encryption starts from the beginning of the alphabet.


Example:


'A' becomes 'N' (a shift of 13 positions)
'B' becomes 'O'
'C' becomes 'P'
...
'M' becomes 'Z' (end of the alphabet, start at the beginning).
'N' becomes 'A' (and so on).
using ROT13:

"HELLO" becomes "URYYB".
"URYYB" (backwards) becomes "HELLO".
ROT13 does not affect characters that are not letters, they remain unchanged.


This encryption method is simple and easily reversible. However, it is used more for hiding information rather than for serious data protection.


![ROT3](https://www.google.com/url?sa=i&url=https%3A%2F%2Fru.wikipedia.org%2Fwiki%2F%25D0%25A8%25D0%25B8%25D1%2584%25D1%2580_%25D0%25A6%25D0%25B5%25D0%25B7%25D0%25B0%25D1%2580%25D1%258F&psig=AOvVaw1hiRy5oWxnEjOwD_gYEHHt&ust=1699145397807000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCIDvx4CQqYIDFQAAAAAdAAAAABAJ)


---


## Reverse or Atbash


Atbash is an encryption method that relies on reversing the order of characters in the alphabet. It is named after the first two letters of the alphabet, "A" and "B", which are reversed.


### What Atbash is and how it works


The alphabet and the numeric values of the letters: First, each letter of the alphabet is assigned a numeric value. For example, 'A' has a value of 1, 'B' has a value of 2, 'C' has a value of 3, and so on until 'Z' which has a value of 26.


Reverse order: the Atbash makes a substitution of each letter with its 'pair' from the end of the alphabet. For example, 'A' is replaced by 'Z', 'B' by 'Y', 'C' by 'X' and so on.


Immutability of other characters: All non-letter characters (numbers, punctuation marks, etc.) remain unchanged.


Reversibility: Atbash is also a self-decoding cipher. Applying it to the encrypted text will produce the original text.


Example:


The original message: "HELLO."


Atbash encryption: "SVOOL".


Atbash decryption: "HELLO".


Atbash is another simple encryption method that can be used to hide information, but like ROT13, is not intended to seriously protect data.


### ROT-X with user input encrypt and decrypt

Caesars chypher with ROT which inputs user

### What is ROT or Caesar cypher

ROT is a simple encryption method that relies on shifting each letter in the alphabet by 13 positions. ROT stands for "rotate" and the number 13 is the number of positions by which the shift occurs.

### How ROT-X works

The way ROT-X works is that each letter is replaced by a letter that is X positions forward. If it reaches the end of the alphabet, encryption starts from the beginning of the alphabet.


Example:

X = 13
'A' becomes 'N' (a shift of 13 positions)
'B' becomes 'O'
'C' becomes 'P'
...
'M' becomes 'Z' (end of the alphabet, start at the beginning).
'N' becomes 'A' (and so on).

X = 1
'A' becomes 'B' (a shift of 1 position)
'B' becomes 'C'
'C' becomes 'D'
...
'M' becomes 'N' (end of the alphabet, start at the beginning).
'N' becomes 'O' (and so on).

using ROTX:

This encryption method is simple and easily reversible. However, it is used more for hiding information rather than for serious data protection.

***

## HOW TO USE

```
$ ./cyphertool

Welcome to the Cypher Tool!

Select operation (1/2):
1. Encrypt.
2. Decrypt.
2

Select cypher (1/2/3):
1. ROT13.
2. Reverse.
3. ROTX with user input
2

Enter the message:
zb

Decrypted message using reverse:
ay
```

## WHAT TO DO LIST

- [X] ~~Search teammates~~
- [X] Submit - Enam, Nikita, Oleh
- [X] Understand how to work with Markdown - Enam, Nikita, Oleh
- [X] ~~Choose third cypher~~ Caesar ROT-X - Enam and Nikita
- [X] Do REVERSE/ATBASH - Enam Hossain
- [X] Do ROT13 - Oleh
- [X] Do README.md - Nikita Strekalov
- [X] 1. Greet the user - Nikita
- [X] 2. Allow the user to select the operation - Nikita
- encrypt
- decrypt
- [X] 3. Allow the user to select the encryption type - Nikita
- [X] 4. Allow the user to input the message to encrypt/decrypt - Nikita, Enam
- [X] 5. Output the result of the operation - Nikita, Enam
- [X] 6. If user input is invalid, the program should keep prompting the user to input again, until valid input is provided. - Nikita
- [X] 7. Before validatid the input, it has to be trimmed - Nikita
- (remove whitespaces from the beginning and the end of the input).
- [X] 8. When encrypting or decrypting, ensure that any non-alphabet characters in the message are left unchanged. - Nikita and Enam
- [X] 9. The tool has to support three encryptions:
- rot13, similar to ShiftBy you created for a single rune.
- reverse, similar to ReverseAlphabetValue you created for a single rune.
- One more encryption of your choice.

|    Oleh    |    Enam    |    Nikita    |
|:-----------|:-----------|:-------------|
|      8     |      8     |      13      |

- [X] Do table just to know how to do it

## WTD list

### Enam Hossain
- [X] Encrypt REVERSE
- [X] Decrypt REVERSE
- [X] A - Z
- [X] a - z
- [X] input
- [X] output 
- [X] non-alpabetical inputs
- [X] unchanged non alhabetical

### Oleh Ilchuk
- [X] Encrypt ROT13
- [X] Decrypt ROT13
- [X] A - Z
- [X] a - z
- [X] input
- [X] output 
- [X] non-alpabetical inputs
- [X] unchanged non alhabetical

### Nikita Strekalov
- [X] Encrypt ROTX (with user input)
- [X] Decrypt ROTX
- [X] A - Z
- [X] a - z
- [X] input
- [X] output 
- [X] non-alpabetical inputs
- [X] unchanged non alhabetical
- [X] greeting
- [X] select operations
- [X] select encryption types
- [X] output
- [X] README.md
