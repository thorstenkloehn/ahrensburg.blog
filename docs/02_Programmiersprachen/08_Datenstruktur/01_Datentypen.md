## stdint.h  Moderne C und C++

#### Ganzzahlige Typen
|Art|Beschreibung|
|---|------------|
|uint8_t| 8-Bit-Ganzzahlen ohne Vorzeichen (0 bis 255)|
|uint16_t|Vorzeichenlose 16-Bit-Ganzzahlen (0 bis 65535)|
|uint32_t|32-Bit-Ganzzahlen ohne Vorzeichen (0 bis 4294967295)|
|uint64_t|64-Bit-Ganzzahlen ohne Vorzeichen (0 bis 18446744073709551615)|
|int8_t|Vorzeichenbehaftete 8-Bit-Ganzzahlen (-128 bis 127)|
|int16_t|Vorzeichenbehaftete 16-Bit-Ganzzahlen (-32768 bis 32767)|
|int32_t|Vorzeichenbehaftete 32-Bit-Ganzzahlen (-2147483648 bis 2147483647)|
|int64_t|Vorzeichenbehaftete 64-Bit-Ganzzahlen (-9223372036854775808 bis 9223372036854775807)|



##  C Datentypen

### Grundtypen

#### Ganzzahlige Typen

Art |Speichergröße|Wertebereich|
|--- | ------------ | -----------|
char|	1 Byte|	-128 bis 127 oder 0 to 255|
unsigned char|	1 Byte|	0 bis 255|
signed char|	1 Byte|	-128 bis 127|
int|	2 oder 4 Bytes|	-32,768 bis 32,767 oder -2,147,483,648 bis 2,147,483,647|
unsigned int|	2 oder 4 Bytes|	0 bis 65,535 oder 0 to 4,294,967,295|
short|	2 Bytes|	-32,768 bis 32,767|
unsigned short|	2 Bytes	| 0 bis 65,535|
long|	8 bytes or (4 Bytes für 32 Bit OS)|	-9223372036854775808 bis 9223372036854775807|
unsigned long|	8 bytes |	0 bis 18446744073709551615|

#### Gleitkommatypen

|Art|Speichergröße|Wertebereich|Präzision|
|---|-------------|------------|---------|
|float|	4 Byte	1.2E-38 bis 3.4E+38|6 Nachkommastellen|
|double|	8 Byte	2.3E-308 bis 1.7E+308 |15 Nachkommastellen|
|long double|	10 Byte	3.4E-4932 bis 1.1E+4932	|19 Nachkommastellen|

#### Der leere Typ

|Art|Präzisio|
|---|--------|
|void|Leere Type|

#### Zeichen

|Art|Präzisio|
|---|-------|
|char|Zeichen|


## C++ Datentypen


### Grundtypen
#### Ganzzahlige Typen

|Art|Typische Bitbreite|Typischer Bereich|
|---|------------------|-----------------|
|char|1 Byte|-127 bis 127 oder 0 bis 255 |
|unsigned char|	1 Byte|	0 bis 255|
|signed char| 1 Byte |-127 bis 127|
|int|4 Bytes|-2147483648 bis 2147483647|
|unsigned int|4 Bytes|0 bis 4294967295|
|signed int|4 Bytes|-2147483648 bis 2147483647|
|short int|	2 Bytes|-32768 bis 32767|
|unsigned short int|2 Bytes|0 bis 65,535|
|signed short int|2 Bytes|-32768 bis 32767|
|long int|8 Bytes|-2,147,483,648 bis 2,147,483,647|
|signed long int|8 Bytes|das gleiche wie long int|
|unsigned long int|	8 Bytes|0 bis 4,294,967,295|
|long long int| 8 Bytes|	-(2^63) bis (2^63)-1|
|unsigned long long int|	8 Bytes|	0 bis 18,446,744,073,709,551,615|

#### Gleitkommatypen

|Art|Typische Bitbreite|
|---|------------------|
|float|4 Bytes|
|double|8 Bytes|
|long double|12 Bytes|	

#### Der leere Typ

|Art|Präzisio|
|---|--------|
|void|Leere Type|

#### Zeichen

|Art|Typische Bitbreite|Typischer Bereich|
|---|------------------|-----------------|
|char|1 Byte|Zeichen|
|wchar_t|2 oder 4 Bytes|Zeichen|

#### Boolean
|Art|Typische Bitbreite|Typischer Bereich|
|---|------------------|-----------------|
|bool| 1 Bit|Ja oder nein|


## Rust Datenypen
### Grundtypen

#### Ganzzahlige Typen

|Größe|Unterzeichnet|Ohne Vorzeichen|
|-----|-------------|---------------|
8 Bit|i8|u8|
16Bit|i16|u16|
32 Bit|i32|u32|
64 Bit|i64|u64|
128 Bit|i128|u128|

#### Gleitkommatypen

|Art|Repräsentiert|
|---|-------------|
|f64|64-Bit-Gleitkommatyp mit doppelter Genauigkeit|
|f32|32-Bit-Gleitkommatyp mit einfacher Genauigkeit|

#### Zeichen

|Art|Präzisio|
|---|-------|
|char|Zeichen|

#### Boolean
|Art|Typische Bitbreite|Typischer Bereich|
|---|------------------|-----------------|
|bool| 1 Bit|Ja oder nein|




## Go Datenypen

### Grundtypen
#### Ganzzahlige Typen
|Art|Beschreibung|
|---|------------|
|uint8| 8-Bit-Ganzzahlen ohne Vorzeichen (0 bis 255)|
|uint16|Vorzeichenlose 16-Bit-Ganzzahlen (0 bis 65535)|
|uint32|32-Bit-Ganzzahlen ohne Vorzeichen (0 bis 4294967295)|
|uint64|64-Bit-Ganzzahlen ohne Vorzeichen (0 bis 18446744073709551615)|
|int8|Vorzeichenbehaftete 8-Bit-Ganzzahlen (-128 bis 127)|
|int16|Vorzeichenbehaftete 16-Bit-Ganzzahlen (-32768 bis 32767)|
|int32|Vorzeichenbehaftete 32-Bit-Ganzzahlen (-2147483648 bis 2147483647)|
|int64|Vorzeichenbehaftete 64-Bit-Ganzzahlen (-9223372036854775808 bis 9223372036854775807)|
#### Gleitkommatypen

|Art|Repräsentiert|
|---|-------------|
|float32|IEEE-754 32-Bit-Gleitkommazahlen|
|float64|IEEE-754 64-Bit-Gleitkommazahlen|
|complex64|Komplexe Zahlen mit float32 Real- und Imaginärteilen|
|Komplex128|Komplexe Zahlen mit float64 Real- und Imaginärteilen|

#### Zeichen

|Art|Präzisio|
|---|-------|
|byte|Zeichen|

#### Boolean
|Art|Typische Bitbreite|Typischer Bereich|
|---|------------------|-----------------|
|bool| 1 Bit|Ja oder nein|
