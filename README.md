# SigSci Agent Report
Reports on the SigSci Agents Across Sites in a Corp 

### Prerequisites

1. sigsciAgentRpt uses an image library that requires the SigSci logo image to be in the directory with the binary. It is bundled.
2. Make sure you have the relevant information in the file config.json. This file should be in the directory along side the binary. The file config.json.example accompanies the build and can be copied or renamed to config.json. The following fields are required:
3. A valid email address that is the Signal Sciences userid. 
4. A valid API token is required for the corresponding email used to authenticate.
5. The Signal Sciences Corp for which to run the report where the email and the token are defined.

### Instructions

1. Download the lastest archive from https://github.com/raymond-sigsci/sigsciAgentRpt.
2. The archives are in the folders linuxBuild, windowsBuild and macosBuild. The archive contains the binary, image file and the example json config.
3. Unzip, edit and rename the config file.
4. Run the command `./sigsciAgentRpt`; `./sigsciAgentRpt.lx`; `sigsciAgentRpt.exe`

The report hits the sites endpoint of the SigSci API and returns the sites for the supplied corp. The report then hits the agents endpoint for each site retireved. The report shows the the site, agent name, agent version and whether the agent is online or offline. The agent name field is a hyperlink to the Signal Sciences Dashboard Agent page. Clicking on the link may require further authentication.

