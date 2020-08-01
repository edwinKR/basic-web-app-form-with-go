console.log("===index.js served===")

const applicantsTable = document.querySelector("table")

fetch("/status")
    .then(response => {
        console.log("response>>>>>", response)
        return response.json()
    })
    .then(allApplicants => {
        console.log("allApplicants====", allApplicants)
        allApplicants.forEach(applicant => {
            const newRow = document.createElement("tr")

            const lastName = document.createElement("td")
            lastName.innerHTML = applicant.lname
            const firstName = document.createElement("td")
            firstName.innerHTML = applicant.fname
            const contact = document.createElement("td")
            contact.innerHTML = applicant.contact

            newRow.appendChild(lastName)
            newRow.appendChild(firstName)
            newRow.appendChild(contact)

            applicantsTable.appendChild(newRow)
        })
    })
