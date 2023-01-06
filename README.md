# Traductions

* [En](https://github.com/Carht/fileHasher/edit/main/README.md#documentation)
* [Es](https://github.com/Carht/fileHasher/edit/main/README.md#documentaci%C3%B3n)

# Documentation

## fileHasher
Return the hash md5 or sha512 of the files in a directory or recursively.

### Build/Compilation

```bash
$ git clone https://github.com/carht/fileHasher
$ cd fileHasher
$ go build .
```

### Usage

This binary implement flags in the same style of unix commands, by default return the
list of the local directory with the hash md5.

#### Flags

* **-h [md5|sha512]** : Active the output hash sum for the files, "md5" or "sha512". Default value **"md5"**
* **-w** : Active the walker mode (recursive).
* **-p filepath** : List the files of the "filepath". Default value **"."**

See [Examples](https://github.com/Carht/fileHasher/edit/main/README.md#ejemplosexamples)

# Documentación

## fileHasher
Lista los archivos de un directorio y retorna su hash md5 o el hash sha512, también es 
recursivo sobre el árbol de directorios.

### Construcción/Compilación

```bash
$ git clone https://github.com/carht/fileHasher
$ cd fileHasher
$ go build .
```

### Uso

Tiene varias "flags" al estilo comandos de unix-like, por defecto retorna la lista de archivos
con su hash md5.

#### Flags:

* **-h [md5|sha512]** : Activa el hash de salida "md5" o "sha512". Valor por defecto **md5**.
* **-w** : Activa el modo recursivo o "walker".
* **-p filepath** : Lista los archivos de dicho directorio. Valor por defecto **"."**

### Ejemplos/Examples:

```bash
$ ./fileHasher
```
![filehasher0](https://user-images.githubusercontent.com/110330581/211024604-f5d92ae5-91c7-45f3-beea-98fc5bf4538a.png)

-----
```
$ ./filehasher -h sha512
```
![filehasher1](https://user-images.githubusercontent.com/110330581/211025560-c371f44e-8a60-494c-b9da-729b6c12d157.png)

-----
```
$ ./filehasher -w
```
![filehasher2](https://user-images.githubusercontent.com/110330581/211025257-b0fc7433-dfd8-4df5-8655-7120c29c5a0a.png)

Retorna una larga lista y recursiva.

Return a long list and recursively.

-----
```
$ ./filehasher -w -p ../gomisc -h sha512
```
![filehasher3](https://user-images.githubusercontent.com/110330581/211026009-422c03bd-ecd1-4461-bec0-1445a6f80a5b.png)

Retorna una larga lista combinando las "flags" de entrada.

Return a long list mix the input flags.
