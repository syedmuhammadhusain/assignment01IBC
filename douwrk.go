//Package greetings shows the greetings
package greetings
//GreetingsString is a global variable
var GreetingsString = "Hello World"
//PrintGreetings is a global function
func PrintGreetings(name string) string {
return GreetingsString + "-" + name

}
