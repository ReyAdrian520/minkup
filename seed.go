package main

import (
    "gorm.io/gorm"
)

func seed(db *gorm.DB) {
    users := []User{
        {
            FirstName: "John",
            Age:       25,
            Interests: []string{"Hiking", "Photography", "Coding"},
            Bio:       "I'm a software engineer who loves to explore nature and capture beautiful moments.",
            Drink:     "Coffee",
            Picture:   "https://randomuser.me/api/portraits/men/1.jpg",
        },
        {
            FirstName: "Emma",
            Age:       22,
            Interests: []string{"Yoga", "Reading", "Traveling"},
            Bio:       "I'm a student who enjoys practicing yoga and reading books. I also love to travel and explore new cultures.",
            Drink:     "Tea",
            Picture:   "https://randomuser.me/api/portraits/women/2.jpg",
        },
        {
			FirstName: "Alice",
			Age:       25,
			Interests: []string{"hiking", "reading", "music"},
			Bio:       "I'm an adventurous person who loves the outdoors.",
			Drink:     "coffee",
			Picture:   "https://randomuser.me/api/portraits/women/1.jpg",
		},
		{
			FirstName: "Bob",
			Age:       30,
			Interests: []string{"cooking", "traveling", "movies"},
			Bio:       "I'm a foodie who loves to explore new cuisines.",
			Drink:     "beer",
			Picture:   "https://randomuser.me/api/portraits/men/2.jpg",
		},
		{
			FirstName: "John",
			Age:       25,
			Interests: []string{"music", "sports", "reading"},
			Bio:       "I'm a music lover and a sports enthusiast.",
			Drink:     "Beer",
			Picture:   "https://picsum.photos/id/1/200/300",
		},
		{
			FirstName: "Jane",
			Age:       28,
			Interests: []string{"movies", "traveling", "hiking"},
			Bio:       "I love watching movies and exploring new places.",
			Drink:     "Wine",
			Picture:   "https://picsum.photos/id/2/200/300",
		},
		{
			FirstName: "Bob",
			Age:       30,
			Interests: []string{"cooking", "music", "photography"},
			Bio:       "I'm a foodie and a music enthusiast.",
			Drink:     "Whiskey",
			Picture:   "https://picsum.photos/id/3/200/300",
		},
		{
			FirstName: "Alice",
			Age:       27,
			Interests: []string{"reading", "yoga", "traveling"},
			Bio:       "I love reading books and practicing yoga.",
			Drink:     "Tea",
			Picture:   "https://picsum.photos/id/4/200/300",
		},
		{
			FirstName: "David",
			Age:       26,
			Interests: []string{"sports", "movies", "traveling"},
			Bio:       "I'm a sports fan and a movie buff.",
			Drink:     "Beer",
			Picture:   "https://picsum.photos/id/5/200/300",
		},
		{
			FirstName: "Sarah",
			Age:       29,
			Interests: []string{"music", "yoga", "hiking"},
			Bio:       "I'm a music lover and a nature enthusiast.",
			Drink:     "Wine",
			Picture:   "https://picsum.photos/id/6/200/300",
		},
		{
			FirstName: "Tom",
			Age:       28,
			Interests: []string{"photography", "traveling", "reading"},
			Bio:       "I love taking photos and exploring new places.",
			Drink:     "Whiskey",
			Picture:   "https://picsum.photos/id/7/200/300",
		},
    }

    for _, user := range users {
        db.Create(&user)
    }
}