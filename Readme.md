# Unpuzzled
A user-first CLI library. 

With the ability to import variables from many different sources: command line flags, environment variables, and configuration variables, it's often not clear what values are set, or why.

Unpuzzled gives you and your users a clear explanation of where variables are being set, and which ones are being overwritten.

Clarity prevents confusion.

# Features
* First class importing from:
    * Environment Variables
    * JSON files
    * TOML files
    * CLI Flags
* Ability to choose the order of precendece (ex. cli flags > JSON > TOML > ENV)
* Main Command and Subcommands
* Defaults to verbose output out of the box, with the ability to turn it off. (app.Silent = true`)
* Warnings on overrides 
    * Warnings on overrides by order of precedence. 
        * If a variable is set as an ENV variable, and a CLI flag overwrites it, let the user know.
        * If two or more variables have references to the same pointer, let the user know.
* Displays what variables are set, and what they are set to.
* Add default values for variables.
* Ability to set Variables as Required.
    * If a value isn't set, print a warning to stdout, and exit.
    * If a variable has a `Default` value, it can never be marked as required, because a valid value will be set.

Left to do:
* More variable types

#### Types of Outputs
##### Missing Required Variables:
Unpuzzled will parse all the inputs, and then list all of the missing required variables before exiting the program. This includes required variables in parent commands.
![required variables](https://github.com/timjchin/unpuzzled/raw/master/fixtures/missing_required_variables.jpg "Required Variable Example CLI Output.")

##### Set Variables
Set Variables can be shown in two outputs.

Stdout is the default:
![set variable stdout](https://github.com/timjchin/unpuzzled/raw/master/fixtures/set_variables_stdout.jpg "Example Stdout Output for set variables.")

And a table option can be chosen by changing `OverridesOutputInTable` to `true`:
```
app := unpuzzled.NewApp()
app.OverridesOutputInTable = true
```
![set variable table view](https://github.com/timjchin/unpuzzled/raw/master/fixtures/set_variables_table_output.jpg "Example Table Output for set variables.")

##### Overwritten Destination
Since unpuzzled uses pointers to set the final values, it's possible that the same pointer may be left in multiple variables. 

Unpuzzled will warn you that these values have been overwritten. 

In the example below, the variables `example-a` and `example-b` both point to `testString`. If both are set, `example-b` will override `example-a`, because it is later in the Variables array. 

```
var testString string
app := unpuzzled.NewApp()
app.Command = &unpuzzled.Command{
    Variables: []unpuzzled.Variable{
        &unpuzzled.StringVariable{
            Name:        "example-a",
            Destination: &testString,
        },
        &unpuzzled.StringVariable{
            Name:        "example-b",
            Destination: &testString,
        },
    },
    Action: func() {
        fmt.Println(testString)
    },
}
app.Run(os.Args)
```
![overwritten pointer](https://github.com/timjchin/unpuzzled/raw/master/fixtures/overwritten_pointer.jpg "Example Output for overwritten variables.")

(Full example in [example/example_ovewritten_pointer.go](https://github.com/timjchin/unpuzzled/blob/master/example/example_overwritten_pointer.go))


#### How to use JSON / Toml configs:
##### TOML:
```
app := unpuzzled.NewApp()
app.Command = &unpuzzled.Command{
    Name: "main",
    Variables: []unpuzzled.Variable{
        &unpuzzled.ConfigVariable{
            StringVariable: &unpuzzled.StringVariable{
                Name: "config"
                Description: "Main configuration, use with `go run main.go --config=path_to_file.toml`",
                Type: unpuzzled.TomlConfig,
            },
        },
    },
}
```
##### JSON Config Example:
```
app := unpuzzled.NewApp()
app.Command = &unpuzzled.Command{
    Name: "main",
    Variables: []unpuzzled.Variable{
        &unpuzzled.ConfigVariable{
            StringVariable: &unpuzzled.StringVariable{
                Name: "config"
                Description: "Main configuration flag, use with `go run main.go --config=path_to_file.json`",
                Type: unpuzzled.JsonConfig,
            },
        },
    },
}
```

Status:
Alpha.