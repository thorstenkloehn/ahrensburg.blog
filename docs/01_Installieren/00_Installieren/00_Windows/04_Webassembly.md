
## Webassembly
### Emscripten
#### Einführung

* Windows 10
* CLion 2019.1 oder höher Version.
* Emscripten
* Chrome

#### Vorbereitung

#### Emscripten SDK

##### Installieren Sie das Emscripten SDK

```

git clone https://github.com/emscripten-core/emsdk.git
cd emsdk
emsdk install latest
emsdk activate latest

```

##### Clion Einstellung.

###### PATH Variables

* Öffnet File -->> Settings -->> PATH Variables
* Erstellen Sie eine Variable EMSCRIPTEN.

##### Konfigurieren Sie die Build-Einstellungen für Emscripten

* Öffnet Sie File -->> Settings -->> Build, Execution, Deployment -->Cmake
* CMake options eingeben : 
```
-DCMAKE_TOOLCHAIN_FILE=${EMSCRIPTEN}/upstream/emscripten/cmake/Modules/Platform/Emscripten.cmake

```
