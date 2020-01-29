package main

import (
    "encoding/json"
    "flag"
//  "fmt"
    "os"
    "github.com/jung-kurt/gofpdf"
    sigsci "github.com/signalsciences/go-sigsci"
)
type Configuration struct {
    Email       string  `json:"email"`
    APItoken    string  `json:"token"`
    Corp        string  `json:"corp"`
}
const (
	apiURL            = "https://dashboard.signalsciences.net/api/v0"
	loginEndpoint     = apiURL + "/auth/login"
    interval          = 300
)

func main() {    
    err := GeneratePdf("sigsciAgentRpt.pdf")
    if err != nil {
        panic(err)
    }
}

// GeneratePdf generates our pdf by adding text and images to the page
// then saving it to a file (name specified in params).
func GeneratePdf(filename string) error {

    c := flag.String("c", "config.json", "Specify the configuration file.")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}
 // For testing Config Structure   fmt.Printf("%+v\n", &Config)

    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 16)

    // ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
    pdf.ImageOptions(
        "sigsci-logo__primary_sm.png",
        0, 0,
        0, 0,
        false,
        gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true},
        0,
        "",
    )

    // CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
    title := "SigSci Agent Report for Corp: "; title += Config.Corp
    pdf.CellFormat(190, 7, "                   ", "0", 1, "CM", false, 0, "")
    pdf.CellFormat(190, 7, title, "0", 1, "CM", false, 0, "")
    pdf.CellFormat(190, 7, "                   ", "0", 1, "CM", false, 0, "")
    pdf.SetFont("Arial", "", 10)
    pdf.CellFormat(40, 7, "Site", "BR", 0, "LM", false, 0, "")
    pdf.CellFormat(95, 7, "Agent", "BR", 0, "LM", false, 0, "")
    pdf.CellFormat(20, 7, "Version", "BR", 0, "LM", false, 0, "")
    pdf.CellFormat(20, 7, "Status", "BR", 1, "LM", false, 0, "")
    pdf.SetFont("Arial", "", 8)

    sc := sigsci.NewTokenClient(Config.Email, Config.APItoken)
    mySites, err := sc.ListSites(Config.Corp)
    if err != nil {
        panic(err)
    }
    var prevSiteName = ` `
    var agentLink = `https://dashboard.signalsciences.net/corps/`
    for _, elem := range mySites {
        agents, err1 := sc.ListAgents(Config.Corp, elem.Name)
        if err1 != nil {
            panic(err1)
        }
        for _, elem1 := range agents {
            if elem.Name != prevSiteName {
                pdf.CellFormat(40, 7, elem.Name, "0", 0, "LM", false, 0, "")
                prevSiteName = elem.Name
                
            } else {
                pdf.CellFormat(40, 7, "             ", "0", 0, "LM", false, 0, "")
            }
            agentLink += Config.Corp; agentLink += `/sites/`; agentLink += elem.Name;
            agentLink += `/agents/`; agentLink += elem1.AgentName
            
            pdf.CellFormat(95, 7, elem1.AgentName, "0", 0, "LM", false, 0, agentLink)
            pdf.CellFormat(20, 7, elem1.AgentVersion, "0", 0, "LM", false, 0, "")
            pdf.CellFormat(20, 7, elem1.AgentStatus, "0", 1, "LM", false, 0, "")

            agentLink = `https://dashboard.signalsciences.net/corps/`
        }
    }   
    
    return pdf.OutputFileAndClose(filename)
}
