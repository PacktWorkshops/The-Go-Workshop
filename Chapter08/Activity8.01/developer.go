package payroll
import (
	"errors"
	"fmt"
)
type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}
func (d Developer) Pay() (string, float64) {
	fullName := d.fullName()
	return fullName, d.HourlyRate * d.HoursWorkedInYear
}
func (d Developer) fullName() string {
	fullName := d.Individual.FirstName + " " + d.Individual.LastName
	return fullName
}
func (d Developer) ReviewRating() error {
	total := 0
	for _, v := range d.Review {
		rating, err := overallReview(v)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n", d.fullName(), averageRating)
	return nil
}
// overallReview determines if the interview concrete type, if it is a string
// we need to convert the string to the corresponding int.
func overallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := convertReviewToInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil

	default:
		return 0, errors.New("unknown type")
	}
}
// convertReviewToInt maps string value to a numeric value
func convertReviewToInt(str string) (int, error) {
	switch str {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("invalid rating: " + str)
	}
}