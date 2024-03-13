package services

import "github.com/harish876/go-templ-htmx-exploration/models"

var Data = models.GridData{
	{
		Id:       "authurmelo",
		Name:     "Arthur Melo",
		Status:   "inactive",
		Position: "Design Director",
		Email:    "authurmelo@example.com",
		Badges:   []string{"Design", "Product", "Marketing"},
		Img:      "https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=880&q=80",
	},
	{
		Id:       "ameliaanderson",
		Name:     "Amelia. Anderson",
		Status:   "active",
		Position: "Lead Developer",
		Email:    "ameliaanderson@example.com",
		Badges:   []string{"Design", "Product", "Marketing"},
		Img:      "https://images.unsplash.com/photo-1531590878845-12627191e687?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=764&q=80",
	},
	{
		Id:       "junior",
		Name:     "junior REIS",
		Status:   "active",
		Position: "Product Manager",
		Email:    "junior@example.com",
		Badges:   []string{"Design", "Product", "Marketing"},
		Img:      "https://images.unsplash.com/photo-1608174386344-80898cec6beb?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=687&q=80",
	},
	{
		Id:       "oliviawathan",
		Name:     "olivia Wathan",
		Status:   "active",
		Position: "Lead Designer",
		Email:    "junior@example.com",
		Badges:   []string{"Design"},
		Img:      "https://images.unsplash.com/photo-1488508872907-592763824245?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80",
	},
	{
		Id:       "mia",
		Name:     "Mia",
		Status:   "active",
		Position: "Graphic Designer",
		Email:    "mia@example.com",
		Badges:   []string{"Design"},
		Img:      "https://images.unsplash.com/photo-1499470932971-a90681ce8530?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80",
	},
	{
		Id:       "johnsmith",
		Name:     "John Smith",
		Status:   "Active",
		Position: "Software Engineer",
		Email:    "john@example.com",
		Badges:   []string{"Engineering"},
		Img:      "https://randomuser.me/api/portraits/men/1.jpg",
	},
	{
		Id:       "emilyjones",
		Name:     "Emily Jones",
		Status:   "Active",
		Position: "UX Designer",
		Email:    "emily@example.com",
		Badges:   []string{"Design"},
		Img:      "https://randomuser.me/api/portraits/women/1.jpg",
	},
	{Id: "davidbrown",
		Name:     "David Brown",
		Status:   "Active",
		Position: "Product Manager",
		Email:    "david@example.com",
		Badges:   []string{"Product"},
		Img:      "https://randomuser.me/api/portraits/men/2.jpg",
	},
	{Id: "sarahwilson",
		Name:     "Sarah Wilson",
		Status:   "Active",
		Position: "Marketing Specialist",
		Email:    "sarah@example.com",
		Badges:   []string{"Marketing"},
		Img:      "https://randomuser.me/api/portraits/women/2.jpg",
	},
	{Id: "michaelclark",
		Name:     "Michael Clark",
		Status:   "Active",
		Position: "Data Analyst",
		Email:    "michael@example.com",
		Badges:   []string{"Analytics"},
		Img:      "https://randomuser.me/api/portraits/men/3.jpg",
	},
	{Id: "laurasmith",
		Name:     "Laura Smith",
		Status:   "Active",
		Position: "Frontend Developer",
		Email:    "laura@example.com",
		Badges:   []string{"Engineering"},
		Img:      "https://randomuser.me/api/portraits/women/3.jpg",
	},
	{Id: "samuelwilson",
		Name:     "Samuel Wilson",
		Status:   "Active",
		Position: "Graphic Designer",
		Email:    "samuel@example.com",
		Badges:   []string{"Design"},
		Img:      "https://randomuser.me/api/portraits/men/4.jpg",
	},
	{Id: "jessicajones",
		Name:     "Jessica Jones",
		Status:   "Active",
		Position: "Content Writer",
		Email:    "jessica@example.com",
		Badges:   []string{"Marketing"},
		Img:      "https://randomuser.me/api/portraits/women/4.jpg",
	},
	{Id: "andrewdavis",
		Name:     "Andrew Davis",
		Status:   "Active",
		Position: "DevOps Engineer",
		Email:    "andrew@example.com",
		Badges:   []string{"Engineering"},
		Img:      "https://randomuser.me/api/portraits/men/5.jpg",
	},
	{Id: "elinawilliams",
		Name:     "Elina Williams",
		Status:   "Active",
		Position: "UI/UX Designer",
		Email:    "elina@example.com",
		Badges:   []string{"Design", "Product", "Marketing"},
		Img:      "https://randomuser.me/api/portraits/women/5.jpg",
	},
}

func FilterById(data []models.GridDataRow, id string) models.GridDataRow {
	for _, row := range data {
		if row.Id == id {
			return row
		}
	}
	return models.GridDataRow{}
}

func UpdateById(data []models.GridDataRow, id string, newData models.GridDataRow) models.GridDataRow {
	for i, row := range data {
		if row.Id == id {
			data[i].Status = newData.Status
			data[i].Position = newData.Position
			return data[i]
		}
	}
	return models.GridDataRow{}
}

func DeleteById(id string) {
	var updatedData []models.GridDataRow
	for _, row := range Data {
		if row.Id != id {
			updatedData = append(updatedData, row)
		}
	}

	Data = updatedData

}
