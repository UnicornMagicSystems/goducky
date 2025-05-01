package main

import (
	"database/sql"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	_ "github.com/marcboeker/go-duckdb"
)

// Person represents a person record
type Person struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

const dbPath = "people.duckdb"

var db *sql.DB

func main() {
	// Initialize database
	initDB()
	defer db.Close()

	// Create a new Fyne application
	a := app.New()
	w := a.NewWindow("Person Database")
	w.Resize(fyne.NewSize(600, 400))

	// Create form widgets
	firstNameEntry := widget.NewEntry()
	firstNameEntry.SetPlaceHolder("Enter First Name")

	lastNameEntry := widget.NewEntry()
	lastNameEntry.SetPlaceHolder("Enter Last Name")

	ageEntry := widget.NewEntry()
	ageEntry.SetPlaceHolder("Enter Age")

	// Load people data
	people := loadPeople()

	// Create a table to display data
	table := widget.NewTable(
		func() (int, int) {
			return len(people) + 1, 4 // +1 for header row
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
				if row < len(people) {
					person := people[row]
					switch i.Col {
					case 0:
						label.SetText(strconv.Itoa(person.ID))
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
	var selectedID int = -1

	// Handle table selection
	table.OnSelected = func(id widget.TableCellID) {
		if id.Row > 0 { // Skip header row
			idx := id.Row - 1 // Adjust for header
			if idx < len(people) {
				person := people[idx]
				selectedID = person.ID
				firstNameEntry.SetText(person.FirstName)
				lastNameEntry.SetText(person.LastName)
				ageEntry.SetText(strconv.Itoa(person.Age))
			}
		}
	}

	// Function to refresh the table
	refreshTable := func() {
		people = loadPeople()
		table.Refresh()
	}

	// Function to clear form
	clearForm := func() {
		selectedID = -1
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

		// Add to database - no need to specify ID as it will use the sequence
		_, err = db.Exec("INSERT INTO people (first_name, last_name, age) VALUES (?, ?, ?)",
			firstName, lastName, age)
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to add record: %v", err), w)
			return
		}

		// Refresh table and clear form
		refreshTable()
		clearForm()

		dialog.ShowInformation("Success", "Record added successfully", w)
	})

	// Update button
	updateButton := widget.NewButton("Update", func() {
		if selectedID == -1 {
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
		_, err = db.Exec("UPDATE people SET first_name = ?, last_name = ?, age = ? WHERE id = ?",
			firstName, lastName, age, selectedID)
		if err != nil {
			dialog.ShowError(fmt.Errorf("failed to update record: %v", err), w)
			return
		}

		// Refresh table and clear form
		refreshTable()
		clearForm()

		dialog.ShowInformation("Success", "Record updated successfully", w)
	})

	// Delete button
	deleteButton := widget.NewButton("Delete", func() {
		if selectedID == -1 {
			dialog.ShowError(fmt.Errorf("please select a record to delete"), w)
			return
		}

		// Confirm deletion
		dialog.ShowConfirm("Confirm", "Are you sure you want to delete this record?", func(confirmed bool) {
			if confirmed {
				// Delete from database
				_, err := db.Exec("DELETE FROM people WHERE id = ?", selectedID)
				if err != nil {
					dialog.ShowError(fmt.Errorf("failed to delete record: %v", err), w)
					return
				}

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

// Initialize the database
func initDB() {
	var err error
	db, err = sql.Open("duckdb", dbPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to open database: %v", err))
	}

	// Create sequence if it doesn't exist
	_, err = db.Exec(`CREATE SEQUENCE IF NOT EXISTS people_id_seq`)
	if err != nil {
		panic(fmt.Sprintf("Failed to create sequence: %v", err))
	}

	// Create table if it doesn't exist with sequence as default for id
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS people (
			id INTEGER PRIMARY KEY DEFAULT nextval('people_id_seq'),
			first_name VARCHAR NOT NULL,
			last_name VARCHAR NOT NULL,
			age INTEGER NOT NULL
		)
	`)
	if err != nil {
		panic(fmt.Sprintf("Failed to create table: %v", err))
	}
}

// Load people from database
func loadPeople() []Person {
	rows, err := db.Query("SELECT id, first_name, last_name, age FROM people ORDER BY id")
	if err != nil {
		fmt.Printf("Error querying database: %v\n", err)
		return []Person{}
	}
	defer rows.Close()

	var people []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Age)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			continue
		}
		people = append(people, p)
	}

	return people
}
