package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Person represents a person record
type Person struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// Database represents our simple database
type Database map[string]Person

const dbPath = "people.duckydb"

func main() {
	// Create a new Fyne application
	a := app.New()
	w := a.NewWindow("Person Database")
	w.Resize(fyne.NewSize(600, 400))

	// Load database
	db := loadDatabase()

	// Create form widgets
	firstNameEntry := widget.NewEntry()
	firstNameEntry.SetPlaceHolder("Enter First Name")

	lastNameEntry := widget.NewEntry()
	lastNameEntry.SetPlaceHolder("Enter Last Name")

	ageEntry := widget.NewEntry()
	ageEntry.SetPlaceHolder("Enter Age")

	// Create a table to display data
	table := widget.NewTable(
		func() (int, int) {
			return len(db) + 1, 4 // +1 for header row
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Wide Content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			label := o.(*widget.Label)
			if i.Row == 0 {
				// Header row
				switch i.Col {
				case 0:
					label.SetText("ID")
				case 1:
					label.SetText("First Name")
				case 2:
					label.SetText("Last Name")
				case 3:
					label.SetText("Age")
				}
				label.TextStyle = fyne.TextStyle{Bold: true}
			} else {
				// Data rows
				row := i.Row - 1 // Adjust for header
				keys := getKeys(db)
				if row < len(keys) {
					person := db[keys[row]]
					switch i.Col {
					case 0:
						label.SetText(person.ID)
					case 1:
						label.SetText(person.FirstName)
					case 2:
						label.SetText(person.LastName)
					case 3:
						label.SetText(strconv.Itoa(person.Age))
					}
				}
			}
		},
	)

	// Set column widths
	table.SetColumnWidth(0, 50)
	table.SetColumnWidth(1, 150)
	table.SetColumnWidth(2, 150)
	table.SetColumnWidth(3, 80)

	// Variable to track selected record
	var selectedID string

	// Handle table selection
	table.OnSelected = func(id widget.TableCellID) {
		if id.Row > 0 { // Skip header row
			keys := getKeys(db)
			idx := id.Row - 1 // Adjust for header
			if idx < len(keys) {
				person := db[keys[idx]]
				selectedID = person.ID
				firstNameEntry.SetText(person.FirstName)
				lastNameEntry.SetText(person.LastName)
				ageEntry.SetText(strconv.Itoa(person.Age))
			}
		}
	}

	// Function to refresh the table
	refreshTable := func() {
		table.Refresh()
	}

	// Function to clear form
	clearForm := func() {
		selectedID = ""
		firstNameEntry.SetText("")
		lastNameEntry.SetText("")
		ageEntry.SetText("")
		table.UnselectAll()
	}

	// Add button
	addButton := widget.NewButton("Add", func() {
		firstName := firstNameEntry.Text
		lastName := lastNameEntry.Text
		ageText := ageEntry.Text

		if firstName == "" || lastName == "" || ageText == "" {
			dialog.ShowError(fmt.Errorf("all fields are required"), w)
			return
		}

		age, err := strconv.Atoi(ageText)
		if err != nil {
			dialog.ShowError(fmt.Errorf("age must be a number"), w)
			return
		}

		// Generate a new ID
		newID := strconv.Itoa(len(db) + 1)

		// Add to database
		db[newID] = Person{
			ID:        newID,
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
		}

		// Save database
		saveDatabase(db)

		// Refresh table and clear form
		refreshTable()
		clearForm()

		dialog.ShowInformation("Success", "Record added successfully", w)
	})

	// Update button
	updateButton := widget.NewButton("Update", func() {
		if selectedID == "" {
			dialog.ShowError(fmt.Errorf("please select a record to update"), w)
			return
		}

		firstName := firstNameEntry.Text
		lastName := lastNameEntry.Text
		ageText := ageEntry.Text

		if firstName == "" || lastName == "" || ageText == "" {
			dialog.ShowError(fmt.Errorf("all fields are required"), w)
			return
		}

		age, err := strconv.Atoi(ageText)
		if err != nil {
			dialog.ShowError(fmt.Errorf("age must be a number"), w)
			return
		}

		// Update database
		person := db[selectedID]
		person.FirstName = firstName
		person.LastName = lastName
		person.Age = age
		db[selectedID] = person

		// Save database
		saveDatabase(db)

		// Refresh table and clear form
		refreshTable()
		clearForm()

		dialog.ShowInformation("Success", "Record updated successfully", w)
	})

	// Delete button
	deleteButton := widget.NewButton("Delete", func() {
		if selectedID == "" {
			dialog.ShowError(fmt.Errorf("please select a record to delete"), w)
			return
		}

		// Confirm deletion
		dialog.ShowConfirm("Confirm", "Are you sure you want to delete this record?", func(confirmed bool) {
			if confirmed {
				// Delete from database
				delete(db, selectedID)

				// Save database
				saveDatabase(db)

				// Refresh table and clear form
				refreshTable()
				clearForm()

				dialog.ShowInformation("Success", "Record deleted successfully", w)
			}
		}, w)
	})

	// Clear button
	clearButton := widget.NewButton("Clear", clearForm)

	// Create form layout
	form := container.NewVBox(
		widget.NewLabel("First Name:"),
		firstNameEntry,
		widget.NewLabel("Last Name:"),
		lastNameEntry,
		widget.NewLabel("Age:"),
		ageEntry,
		container.NewHBox(
			layout.NewSpacer(),
			addButton,
			updateButton,
			deleteButton,
			clearButton,
			layout.NewSpacer(),
		),
	)

	// Create main layout
	content := container.NewBorder(
		form,
		nil,
		nil,
		nil,
		container.NewPadded(table),
	)

	w.SetContent(content)
	w.ShowAndRun()
}

// Helper function to get keys from map in a consistent order
func getKeys(db Database) []string {
	keys := make([]string, 0, len(db))
	for k := range db {
		keys = append(keys, k)
	}
	return keys
}

// Load database from file
func loadDatabase() Database {
	db := make(Database)

	// Check if file exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		return db
	}

	// Read file
	data, err := ioutil.ReadFile(dbPath)
	if err != nil {
		fmt.Printf("Error reading database file: %v\n", err)
		return db
	}

	// Unmarshal JSON
	err = json.Unmarshal(data, &db)
	if err != nil {
		fmt.Printf("Error parsing database file: %v\n", err)
		return make(Database)
	}

	return db
}

// Save database to file
func saveDatabase(db Database) {
	// Marshal to JSON
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		fmt.Printf("Error serializing database: %v\n", err)
		return
	}

	// Write to file
	err = ioutil.WriteFile(dbPath, data, 0644)
	if err != nil {
		fmt.Printf("Error writing database file: %v\n", err)
	}
}
