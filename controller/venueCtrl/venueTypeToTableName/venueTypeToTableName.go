package venueTypeToTableName

func VenueTypeToTableName(venueType string) (tableName string) {
	switch venueType {
	case "羽毛球":
		tableName = "badminton_venue_states"
	case "乒乓球":
		tableName = "table_tennis_venue_states"
	case "网球":
		tableName = "tennis_venue_states"
	}
	return tableName
}
