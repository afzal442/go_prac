package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Emp_id     int
	Emp_name   string
	Emp_age    int
	Emp_dep    string
	Emp_sub_id int
}

// var db *sql.DB

func main() {
	// Capture connection properties.
	db, err := sql.Open("mysql", "root:hiclass@12@/db_org2")

	// Get a database handle.
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Create a new record.
	emp_details, err := Emp_details("Ravi", db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Emp_details found: %v\n", emp_details)

	// Hard-code ID 2 here to test the query.
	emp, err := Emp_detailsByID(4, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Employee found: %v %v\n", emp.Emp_name, emp.Emp_age)

	emp_detail, err := addEmployee(Employee{
		Emp_id:   7,
		Emp_name: "Satish",
		Emp_age:  45,
		Emp_dep:  "Management",
	}, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Details of added album: %v\n", emp_detail)

	rowsAffected, err2 := EmpUpdate(db, 2)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Rows Affected:", rowsAffected)
	}

	// rows, err := db.Query("SELECT * FROM album; SELECT * FROM song;")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// // Loop through the first result set.
	// for rows.Next() {
	// 	// Handle result set.
	// }

	// 	// Advance to next result set.
	// 	rows.NextResultSet()

	// 	// Loop through the second result set.
	// 	for rows.Next() {
	// 		// Handle second set.
	// 	}

	// 	// Check for any error in either result set.
	// 	if err := rows.Err(); err != nil {
	// 		log.Fatal(err)
	// 	}
}

// albumsByArtist queries for albums that have the specified artist name.
func Emp_details(emp_name string, db *sql.DB) ([]Employee, error) {
	// An albums slice to hold data from returned rows.
	var emp_details []Employee

	rows, err := db.Query("SELECT * FROM org_kloudone WHERE Emp_name = ?", emp_name)
	if err != nil {
		return nil, fmt.Errorf("Emp_detailsbyname %q: %v", emp_name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var emp_d Employee
		if err := rows.Scan(&emp_d.Emp_id, &emp_d.Emp_name, &emp_d.Emp_age, &emp_d.Emp_dep, &emp_d.Emp_sub_id); err != nil {
			return nil, fmt.Errorf("Emp_detailsbyname %q: %v", emp_name, err)
		}
		emp_details = append(emp_details, emp_d)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Emp_detailsbyname %q: %v", emp_name, err)
	}
	return emp_details, nil
}

// albumByID queries for the album with the specified ID.
func Emp_detailsByID(id int64, db *sql.DB) (Employee, error) {
	// An album to hold data from the returned row.
	var emp Employee

	row := db.QueryRow("SELECT * FROM org_kloudone WHERE Emp_id = ?", id)
	if err := row.Scan(&emp.Emp_id, &emp.Emp_name, &emp.Emp_age, &emp.Emp_dep, &emp.Emp_sub_id); err != nil {
		if err == sql.ErrNoRows {
			return emp, fmt.Errorf("Emp_detailsByID %d: no such employee details", id)
		}
		return emp, fmt.Errorf("Emp_detailsByID %d: %v", id, err)
	}
	return emp, nil
}

// addEmployee adds the specified Employee to the database,
// returning the Employee ID of the new entry
func addEmployee(emp Employee, db *sql.DB) (int, error) {

	result, err := db.Exec("INSERT INTO org_kloudone (Emp_id, Emp_name, Emp_age, Emp_dep) VALUES (?, ?, ?, ?)", emp.Emp_id, emp.Emp_name, emp.Emp_age, emp.Emp_dep)
	if err != nil {
		return 0, fmt.Errorf("addEmployee: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addEmployee: %v %v", id, err)
	}
	return emp.Emp_id, nil
}

func EmpUpdate(db *sql.DB, emp_id int) (int64, error) {
	result, err := db.Exec("update org_kloudone set Emp_name = 'Vijay', Emp_age = 34 where Emp_id = ?", 2)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
