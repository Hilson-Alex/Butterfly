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

For now, to get the compiler you just need to download the executable [here](https://github.com/Hilson-Alex/Butterfly/releases/tag/0.1.0)
This probably will change on later versions with the need to get resources

### Using The Compiler

The Butterfly compiler receive a directory with one or more *.bf* files inside.

To use the compiler you can either add the path to the `Path` environment
variable on windows and then use it like this: 

```shell
butterfly <folder_with_bf_files>
```

> You can see a tutorial about configuring the environment variables [here.](https://www.c-sharpcorner.com/article/how-to-addedit-path-environment-variable-in-windows-11/)

Or, alternatively, you can just specify the path to the executable as follows:

```shell
path\to\butterfly.exe <folder_with_bf_files>
```

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

## Author

- [Hilson A. W. Junior][Hilson-Alex] - *Initial work*

## License

This project is under the BSD 3-Clause License - for more info read the [LICENSE](LICENSE) file

[Hilson-Alex]: https://github.com/Hilson-Alex
