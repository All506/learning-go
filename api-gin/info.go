package main

//Structure definition for each player 

type player struct {
	ID			int `json:"id"`
	Firstname	string `json:"firstname"`
	Lastname	string `json:"lastname"`
	Position 	string `json:"position"`
	Team		team   	`json:"team"`
}

type team struct {
	ID			int	   `json:"id"`
	Name		string `json:"name"`
	Founded		string `json:"founded"`
	Nickname	string `json:"nickname"`
	Country		string `json:"country"`
}

var players = []player{

	{ID: 1, Firstname: "Kevin", Lastname: "De Bruyne", Position: "Midfielder", Team: teams[0]},
	{ID: 2, Firstname: "Raheem", Lastname: "Sterling", Position: "Forward", Team: teams[0]},
	{ID: 3, Firstname: "Aymeric", Lastname: "Laporte", Position: "Defender", Team: teams[0]},
	{ID: 4, Firstname: "Phil", Lastname: "Foden", Position: "Midfielder", Team: teams[0]},
	{ID: 5, Firstname: "Bernardo", Lastname: "Silva", Position: "Midfielder", Team: teams[0]},
	{ID: 6, Firstname: "Ederson", Lastname: "Moraes", Position: "Goalkeeper", Team: teams[0]},
	{ID: 7, Firstname: "Kyle", Lastname: "Walker", Position: "Defender", Team: teams[0]},
	{ID: 8, Firstname: "Fernandinho", Lastname: "Luiz", Position: "Midfielder", Team: teams[0]},
	{ID: 9, Firstname: "Gabriel", Lastname: "Jesus", Position: "Forward", Team: teams[0]},
	{ID: 10, Firstname: "Riyad", Lastname: "Mahrez", Position: "Forward", Team: teams[0]},
	{ID: 11, Firstname: "John", Lastname: "Stones", Position: "Defender", Team: teams[0]},
	{ID: 12, Firstname: "Ilkay", Lastname: "Gundogan", Position: "Midfielder", Team: teams[0]},

	{ID: 13, Firstname: "Bruno", Lastname: "Fernandes", Position: "Midfielder", Team: teams[1]},
	{ID: 14, Firstname: "Marcus", Lastname: "Rashford", Position: "Forward", Team: teams[1]},
	{ID: 15, Firstname: "Harry", Lastname: "Maguire", Position: "Defender", Team: teams[1]},
	{ID: 16, Firstname: "Paul", Lastname: "Pogba", Position: "Midfielder", Team: teams[1]},
	{ID: 17, Firstname: "Edinson", Lastname: "Cavani", Position: "Forward", Team: teams[1]},
	{ID: 18, Firstname: "David", Lastname: "De Gea", Position: "Goalkeeper", Team: teams[1]},
	{ID: 19, Firstname: "Mason", Lastname: "Greenwood", Position: "Forward", Team: teams[1]},
	{ID: 20, Firstname: "Aaron", Lastname: "Wan-Bissaka", Position: "Defender", Team: teams[1]},
	{ID: 21, Firstname: "Scott", Lastname: "McTominay", Position: "Midfielder", Team: teams[1]},
	{ID: 22, Firstname: "Fred", Lastname: "Rodrigues", Position: "Midfielder", Team: teams[1]},
	{ID: 23, Firstname: "Victor", Lastname: "Lindelof", Position: "Defender", Team: teams[1]},
	{ID: 24, Firstname: "Luke", Lastname: "Shaw", Position: "Defender", Team: teams[1]},
}

var teams = []team{
	{ID: 1,Name:"Manchester United",Founded: "24/04/1902",Nickname: "The Red Devils", Country: "United Kingdom"},
	{ID: 2, Name: "Manchester United", Founded: "24/04/1902", Nickname: "Citizens", Country: "United Kingdom"},
}