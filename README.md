# Build CLI AI Agent with golang

how to do it step by step

1 - go mod init

3 - vim .env

4 - create open router api key paste in .env
as OPENROUTER_API_KEY

5 - go get "github.com/openai/openai-go/v2"

6 - vim main.go

7 - defined url for the openrouter api
define api key from openrouter

check if api is not defiend

create the llm client via openai sdk new client
provide it with base url of openrouter and api key

create the messages array

append your query to the messages array
via openai.usermessage

choose a model and provide its key

define chat completion params

define ctx 

call a chat completion via client.
chat.completins.new

provide the context and the params

print the output

export env vars -> `export $(cat .env | xargs)`

run the first version

go run main.go


second version will be creating a loop
waiting for input and respond from that input

first we create a scanner via
`bufio.newscanner()` and we give it the 
stdin which is the standered input

we print a symbol to let us know we are wainting
for user input like `fmt.print(">")`

then we check if we have scan err
`if !scanner.scan()`
we break the loop

we define the input which will be the text
scanned `scanner.text` without the spaces

`strings.trimspace(scanner.text)`

if that is empty we ignore via continue

then if not we append to the params messages array

we then send the chat completion request

we check for err and then print the response

we run version 2 (interactive chat cli agent)

third version:

we add the response of the llm to the messages array

we then print the messages count

we add a check if messages array become too much 
we clean it and get only the last 4 messages

then we contineu the chat

we add color via faith color library

`go get "github.com/fatih/color"`

we add color by wrapping our strings in print
statement with color.colorstring(the string)

this way we can have better look
and we can understand the ai response better
by giving it different color that our query

and other info as well
