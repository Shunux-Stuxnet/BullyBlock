package report

import (
	"fmt"
	"os"
	"strings"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/technoreck/BB_report/server"
)

func report() {

	Name := "User1"
	Age := "23"
	Email := "TestUser1@gmail.com"
	Phone := "9187753428"
	Desc := "I was a victim of cyberbullying a few years ago, and it was one of the most difficult experiences of my life. It all started when I received a friend request on social media from someone I didn't know. I accepted the request, thinking that it was just a new friend who wanted to connect. However, soon after accepting the request, I started receiving harassing messages and comments on my posts. The messages were filled with insults, derogatory comments, and even threats of physical harm. At first, I tried to ignore the messages and block the sender, but they just created new accounts and continued to harass me. As the bullying continued, I started to feel isolated and alone. I was afraid to tell anyone about what was happening, fearing that it would only make things worse. Eventually, I confided in a close friend who encouraged me to seek help. I reported the bullying to the social media platform and to the authorities. While the bullying did eventually stop, the experience left a lasting impact on me. I felt like I couldn't trust anyone online and was constantly looking over my shoulder, wondering if the bullies were still watching me. In retrospect, I wish I had spoken up sooner and sought help from trusted adults. Cyberbullying can be just as damaging as physical bullying, and no one should have to suffer through it alone."
	Img := "images/Case.png"

	var complain = server.Complain{Name: Name, Age: Age, Email: Email, Phone: Phone, Description: Desc, Image: Img}
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
	buildheading(m)
	buildFruitList(m)
	severity := "low"
	sev := strings.Split(complain.Description, "")
	highsev := []string{"Threats", "Rumors", "Exclusion", "Shaming", "Harassment", "Vulgar", "Invasion", "Offensive", "attack", "Threatening", "Humiliation", "Revenge", "porn", "pornography", "Blackmail", "Cyberstalking", "Doxing", "Impersonation", "Flaming", "Intimidation", "Abusive", "Discrimination"}
	midsev := []string{"Trolling", "Name-calling", "Insults", "Demeaning", "Gossip", "Teasing", "Criticism", "Lies", "Hate", "Discriminatory", "Slander", "mobbing", "Body-shaming", "blaming", "Gaslighting", "Coercion", "Deception", "Jealousy", "Disrespect", "Hostility"}

	for i := 0; i < len(sev); i++ {
		if isAvailable(highsev, sev[i]) {
			severity = "High"
			break
		}
	}
	if severity == "low" {
		for i := 0; i < len(sev); i++ {
			if isAvailable(midsev, sev[i]) {
				severity = "Medium"
				break
			}
		}
	}
	content(m, complain.Name, complain.Age, complain.Phone, complain.Email, complain.Description, complain.Image, severity)

	err := m.OutputFileAndClose("pdfs/Bully_report.pdf")

	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}

func buildheading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/BullyBlock.png", props.Rect{
					Center:  true,
					Percent: 75,
				})
				if err != nil {
					fmt.Println("Image file was not found", err)
				}
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("We help you fight against Cyber Bullying and make a healthy cyberspace", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
				Color: color.NewBlack(),
			})
		})
	})
}

func buildFruitList(m pdf.Maroto) {

	m.SetBackgroundColor(getDarkGreenColor())

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Complain Report", props.Text{
				Top:    1.75,
				Size:   12,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

}

func content(m pdf.Maroto, name string, age string, phoneNo string, email string, Desc string, Image string, severity string) {

	m.SetBackgroundColor(color.NewWhite())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Name: "+name, props.Text{
				Top:    5,
				Size:   10,
				Color:  color.NewBlack(),
				Family: consts.Helvetica,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Age: "+age, props.Text{
				Top:    5,
				Size:   10,
				Color:  color.NewBlack(),
				Family: consts.Helvetica,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("E-mail: "+email, props.Text{
				Top:    5,
				Size:   10,
				Color:  color.NewBlack(),
				Family: consts.Helvetica,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Phone Number: "+phoneNo, props.Text{
				Top:    5,
				Size:   10,
				Color:  color.NewBlack(),
				Family: consts.Helvetica,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Severity: "+severity, props.Text{
				Top:    5,
				Size:   10,
				Color:  color.NewBlack(),
				Family: consts.Helvetica,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Description: ", props.Text{
				Top:    5,
				Size:   10,
				Color:  color.NewBlack(),
				Family: consts.Helvetica,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text(Desc, props.Text{
				Top:    1,
				Size:   10,
				Color:  color.NewBlack(),
				Family: consts.Helvetica,
				Style:  consts.Bold,
				Align:  consts.Left,
			})
		})
	})

	// m.Row(10, func() {
	// 	m.Col(12, func() {
	// 		m.Text("Image: ", props.Text{
	// 			Size:   10,
	// 			Color:  color.NewBlack(),
	// 			Family: consts.Helvetica,
	// 			Style:  consts.Bold,
	// 			Align:  consts.Left,
	// 		})
	// 	})
	// })

	// m.Row(50, func() {
	// 	m.Col(12, func() {
	// 		err := m.FileImage(Image, props.Rect{
	// 			Top:     10,
	// 			Center:  true,
	// 			Percent: 75,
	// 		})
	// 		if err != nil {
	// 			fmt.Println("Image file was not found", err)
	// 		}
	// 	})
	// })

}

func getDarkGreenColor() color.Color {
	return color.Color{
		Red:   15,
		Green: 100,
		Blue:  50,
	}
}

func isAvailable(alpha []string, str string) bool {

	// iterate using the for loop
	for i := 0; i < len(alpha); i++ {
		// check
		if alpha[i] == str {
			// return true
			return true
		}
	}
	return false
}
