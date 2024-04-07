package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/icza/s2prot"
	"github.com/icza/s2prot/rep"
)


func printStructProperties(s interface{}) {
    v := reflect.ValueOf(s)
    t := v.Type()

    for i := 0; i < v.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)

        fmt.Printf("%s: %v\n", field.Name, value.Interface())
    }
}

type UnitTypeData interface {
	Data() string
}




func main() {

	// slowerString := "Slower"
	// slowerModifer := 60

	// slowString := "Slow"
	// slowModifier := 45

	// normalString := "Normal"
	normalModifier := 36

	fastString := "Fast"
	fastModifier := 30

	// fasterString := "Faster"
	// fasterModifier := 26

	



	// r, err := rep.NewFromFile("./Oceanborn LE.SC2Replay")

	// r, err := rep.NewFromFile("./TestNormalSpeed.SC2Replay")

	r, err := rep.NewFromFile("./TestFastSpeed.SC2Replay")


	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}



	// for _, p := range r.Details.Players() {
	// 	fmt.Printf("\tName: %-20s, Race: %c, Team: %d, Toon: %s Result: %v\n",
	// 		p.Name, p.Race().Letter, p.TeamID()+1, p.Toon, p.Result())
	// }

	fmt.Printf("Game Speed: %s\n", r.Details.GameSpeed().Name)

	

	// // Initialize variables to track worker count and supply
	workerCounts := make(map[string]int) // Map of player ID to worker count
	supplyCounts := make(map[string]int) // Map of player ID to total supply

	// uniqueTrackerEvents := make(map[string]s2prot.Event)

	 uniqueTrackerEvents := []s2prot.Event{} ;


	for _, player := range r.Details.Players() {
		workerCounts[player.Toon.String()] = 0;
		supplyCounts[player.Toon.String()] = 0;
	}


	for _, evt := range r.TrackerEvts.Evts {


		var unitTypeNameRaw = evt.Struct.Value("unitTypeName")

		if unitTypeNameRaw != nil {

			if s, ok := unitTypeNameRaw.(string); ok {
				// fmt.Println("The value stored in x is:", s)

				if s != "Overlord" {
					continue;
				}

			} else {
				// fmt.Println("The value stored in x is not a string")
			}

		}



// || unitTypeName != "Overlord"
		if (strings.ToLower(evt.EvtType.Name) != "unitborn" ) {
			// fmt.Printf("%s", lower)
			continue;

		}

		if (evt.Loop() == 0) {
			continue;
		}

		uniqueTrackerEvents = append(uniqueTrackerEvents, evt)

		





		// break;


		// printStructProperties(evt);

		// uniqueTrackerEvents[evt.EvtType.Name] = struct{}{};
	}


	sort.Slice(uniqueTrackerEvents, func(i, j int) bool {
		return uniqueTrackerEvents[i].Loop() < uniqueTrackerEvents[j].Loop()
	})

	modification:= 0

	if r.Details.GameSpeed().Name == fastString {
		modification = fastModifier
	}

	for evtIndex := range uniqueTrackerEvents {
		evt:= uniqueTrackerEvents[evtIndex];

		//  fmt.Printf("\tEvtTypeName: %-20s\n", evt.EvtType.Name);

		// // Calculate in-game time in seconds
		// inGameTime := float64(evt.Loop()*125) / float64(2*16)
		//  inGameTime := float64(evt.Loop()* 125 ) / float64(2)
		realTimeValue:= float64((( evt.Loop() * 125 )) / 2)/1000
		// inGameTime:= realTimeValue  * (36/26)

		inGameTime:= realTimeValue * (float64(modification)/ float64(normalModifier))

		var unitTypeNameRaw = evt.Struct.Value("unitTypeName")
		var unitTypeName string;

		if unitTypeNameRaw != nil {

			if s, ok := unitTypeNameRaw.(string); ok {
				// fmt.Println("The value stored in x is:", s)

				unitTypeName = s
			} else {
				// fmt.Println("The value stored in x is not a string")
			}

			fmt.Printf("Unit Type Name: %s\n" , unitTypeName)
		}


		printStructProperties(evt);

		// fmt.Println("Loop:", evt.Loop())

		// Print the in-game time
		fmt.Println("In-game time:", inGameTime, "seconds")
	}

	

	// // Iterate through game events
	// for _, evt := range r.GameEvts {
	// 	// Check if the event is relevant for worker count or supply

	// 	// fmt.Printf("\tName: %-20s, ID %s, TypeName: %s, TypeId: %d\n", evt.Name, evt.ID, evt.EvtType.Name, evt.EvtType.ID)
	// 	uniqueGameEvents[evt.Name] = struct{}{}
	// }

	// for key := range uniqueGameEvents {

	// 	 fmt.Printf("\tName: %-20s\n", key);
	// }



	// for key, value := range workerCounts {

	// 	fmt.Printf("Key: %s, Value: %d", key, value);
	// }




	// fmt.Printf("Read file contents! %v\n", r)

	// fmt.Printf("Full Header:\n%v\n", r.Header)

	// fmt.Printf("Version:        %v\n", r.Header.VersionString())
	// fmt.Printf("Loops:          %d\n", r.Header.Loops())
	// fmt.Printf("Length:         %v\n", r.Header.Duration())
	// fmt.Printf("Map:            %s\n", r.Details.Title())
	// fmt.Printf("Game events:    %d\n", len(r.GameEvts))
	// fmt.Printf("Message events: %d\n", len(r.MessageEvts))
	// fmt.Printf("Tracker events: %d\n", len(r.TrackerEvts.Evts))

	// fmt.Println("Players:")
	// for _, p := range r.Details.Players() {
	// 	fmt.Printf("\tName: %-20s, Race: %c, Team: %d, Result: %v\n",
	// 		p.Name, p.Race().Letter, p.TeamID()+1, p.Result())
	// }
	defer r.Close()

}