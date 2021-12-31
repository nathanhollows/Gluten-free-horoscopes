package handler

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

// HoroscopeHandler serves each horoscope page
func HoroscopeHandler(w http.ResponseWriter, r *http.Request) error {
	data := make(map[string]interface{})
	data["starsign"] = chi.URLParam(r, "starsign")
	data["title"] = strings.Title(chi.URLParam(r, "starsign"))
	data["date"] = time.Now().Format("January 02")
	horoscope := Horoscope{
		Date:     time.Now().Format("January 02"),
		Live:     rand.Intn(5) + 1,
		Laugh:    rand.Intn(5) + 1,
		Love:     rand.Intn(5) + 1,
		Luck:     rand.Intn(5) + 1,
		Bullshit: 5,
	}

	scopes := []string{}
	scopes = append(scopes, "Watch out for the communist agenda.")
	scopes = append(scopes, "You knead to be careful this week. Beware the gluten intolerant intolerant.")
	scopes = append(scopes, "Hot damn. You’re going to be hungover as shit this week. (That’s fine. Spirits have no gluten in them.)")
	scopes = append(scopes, "Soon retro-graded Saturn is gonna be out of fashion, if you’re looking to impress at the work do then go with velvet instead.")
	scopes = append(scopes, "The Star of David is in House of the Rising Sun, today is not the day to eat cheese and crackers.")
	scopes = append(scopes, "You will meet the lover that reaches the open road. They will be wielding a baguette.")
	scopes = append(scopes, "You should try making some pasta this week. Pasta is delicious and quick to make, and will bring you joy in addition to nutrition.")
	scopes = append(scopes, "This week you’ll hold the door for some people and then you’ll feel a little anxious about that. You probably shouldn’t feel anxious, it’s just a door.")
	scopes = append(scopes, "Bread heals all. Too bad for you, eh?")
	scopes = append(scopes, "You will definitely marry that person you’re thinking about. Stop annoying your friends about it.")
	scopes = append(scopes, "Keep stewing vegetables for days in advance knowing that one day you are going to get it right. You know in your heart that it’s going to happen.")
	scopes = append(scopes, "Today the monkey buccaneer is left handed.")
	scopes = append(scopes, "Moon is in Gatorade")
	scopes = append(scopes, "When the moon hits your eye like a big pizza pie that's a bad omen.")
	scopes = append(scopes, "Which zodiac sign do you identify with?")
	horoscope.Text = scopes[rand.Intn(len(scopes))]
	data["horoscope"] = horoscope

	if r.Header.Get("HX-Request") == "true" {
		data["contentOnly"] = true
	}
	return render(w, data, "horoscope.html")
}

type Horoscope struct {
	Date     string
	Text     string
	Live     int
	Laugh    int
	Love     int
	Luck     int
	Bullshit int
}
