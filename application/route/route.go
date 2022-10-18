package route

import "errors"

type Route struct{
	ID string
	ClientID String
	Positions []Position
}

type Position struct {
	Lat float64
	Long float64
}

func(r *Route) LoadPositions() error {

	if r.ID == "" {
		return errors.New("route id not informed.")
	}
	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}

	defer f.Close()

	sanner := buffio.NewScanner(f)

	for scanner.Scan() {
		data := strings.Split(sanner.Text(), ",")
		lat, err := srtconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}
		long, err := srtconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}
		r.Positions = append(r.Positions, Position {
			Lat: lat,
			Long: long
		})
	}
	return nil
}
