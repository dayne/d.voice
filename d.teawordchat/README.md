# d.teawordchat

Code Explanation:
- Initialization: The textinput is centered using placeholders and character limits. The scroll box is just an array of strings that grows as new words are added.
- Event Handling: It listens for "enter" or "space" keypresses to determine when a word is complete. Upon completion, it clears the text input and appends the word to the scroll box.
- View Rendering: The words are displayed above the input field, showing the user their inputs as they type new words.

## Tasks

### Build

```
go build 
```

### Setup

```
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/bubbles/list
go get github.com/charmbracelet/bubbles/textinput
go get github.com/charmbracelet/lipgloss
go get github.com/eclipse/paho.mqtt.golang
```


