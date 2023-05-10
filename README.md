# Butterfly

Butterfly is an experimental programming language with an event-driven perspective to create concurrent programs by 
default. The Butterfly language is based on modules that communicate exclusively through events and asynchronous message 
passing.

<a href="https://drive.google.com/uc?export=view&id=1vI3_UjFAzELrp2uUPGPjE3CK-NpqgJeN">
  <img src="https://drive.google.com/uc?export=view&id=1vI3_UjFAzELrp2uUPGPjE3CK-NpqgJeN" style="width: 650px; max-width: 100%; height: auto"  alt="a Butter Flying"/>
</a>

*Butterfly image by me, [Hilson-Alex]*


## Getting Started

By following these instructions you will get the working compiler on your computer. 


### Installing

The easiest way to install the compiler, you just need to download the bf_installer.msi [here](https://github.com/Hilson-Alex/Butterfly/releases/tag/0.2.0). 
When executed, the installer will install the executable, dependencies and do all setup needed.

Another way to get the compiler working is downloading the source code, adding the path to the [/out] folder(./out) on the `Path`
environment variable an then compile the source code with the `-o ./out/butterfly.exe` flag. It is important to not change the locations 
of the executable or the resource folder (and anything inside it), or the compiler may malfunction.

> You can see a tutorial about configuring the environment variables [here.](https://www.c-sharpcorner.com/article/how-to-addedit-path-environment-variable-in-windows-11/)


### Using The Compiler

The Butterfly compiler receive a directory with one or more *.bf* files inside. You can read the specification of the Butterfly Language 
Grammar [here](./out/resources/doc/Grammar.md)

If you used the installer or added the [/out](./out) folder to the `Path` environment
variable, then you can just use it like this: 

```shell
butterfly <flags> <folder_with_bf_files>
```

Or, alternatively, you can just specify the path to the executable as follows:

```shell
path\to\butterfly.exe <flags> <folder_with_bf_files>
```
> As said before, it's important to keep the compiler together side-by-side with the resources folder
> Otherwise it will break.

All flags are optional. The valid flags are listed bellow.

- *-help*: prints the usage and all valid flags
- *-out*: especify the path and the name where the executable will be generated.
- *-useQueue*: Builds the executable with an embed event queue instead of immediate event share.
- *-listTokens*: list all valid tokens. You also can read them [here](./out/resources/doc/Tokens.md).
  This is useful to understand the parser errors.
- *-lexerVerbose*: Prints the parsed tokens while compiling. Used on debugging.
- *-keepGeneratedFiles*: Keep the generated files after compiling. Used **ONLY** for debugging.


### Cloning and Building

To get the repository running locally on your machine you will need to install the [Go compiler](https://go.dev/dl/)

You can clone the project by running

```shell
git clone git@github.com:Hilson-Alex/Butterfly.git
```

or, if you can't use ssh

```shell
git clone https://github.com/Hilson-Alex/Butterfly.git
```

Next, open the project folder and run

```shell
go generate ./src
go build -o ../out/butterfly.exe github.com/Hilson-Alex/Butterfly/src
```

And then you will have a `butterfly.exe` on the `out` folder, beside the src.

You can setup the `Path` environment variable as instructed [above](#installing) if wanted


## Author

- [Hilson A. W. Junior][Hilson-Alex] - *Initial work*


## License

This project is under the BSD 3-Clause License - for more info read the [LICENSE](LICENSE) file

[Hilson-Alex]: https://github.com/Hilson-Alex
