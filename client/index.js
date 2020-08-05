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

            const applicantId = document.createElement("td")
            applicantId.innerHTML = applicant.customerId
            const lastName = document.createElement("td")
            lastName.innerHTML = applicant.lname
            const firstName = document.createElement("td")
            firstName.innerHTML = applicant.fname
            const contact = document.createElement("td")
            contact.innerHTML = applicant.contact

            const deleteBtn = document.createElement("button")
            deleteBtn.setAttribute("onclick", `deleteAction(${applicant.customerId})`)
            deleteBtn.innerHTML = "Delete"

            newRow.setAttribute("id", applicant.customerId)

            newRow.appendChild(applicantId)
            newRow.appendChild(lastName)
            newRow.appendChild(firstName)
            newRow.appendChild(contact)

            newRow.appendChild(deleteBtn)

            applicantsTable.appendChild(newRow)
        })
    })

function deleteAction(id) {
    const applicantData = {
        uid: id.toString(),
    }

    const deleteMethod = {
        method: "DELETE",
        headers: {
            "Content-type": "application/json; charset=UTF-8"
        },
        body: JSON.stringify(applicantData),
    }
    console.log(deleteMethod)
    fetch("/delete", deleteMethod)
        .then(response => {
            console.log(response)
        })
        .then(data => {
            console.log(data)
            location.reload()
        })
        .catch(err => console.log(err))
}