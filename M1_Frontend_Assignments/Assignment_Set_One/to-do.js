let addBtn = document.getElementById('addBtn');
let mainList = document.getElementById('main-list');
addBtn.addEventListener('click', addAssignment);

function addAssignment(e) {
    if (mainList.children[0].id == "emptyMsg") {
        mainList.children[0].remove()
    }
    let currentBtn = e.currentTarget;
    let currentInput = currentBtn.previousElementSibling
    let currentAssName = currentInput.value

    let newLi = document.createElement('li')
    newLi.className = "list-group-item d-flex justify-content-between"
    newLi.innerHTML = `<h4 class="flex-grow-1">${currentAssName}</h4>
    <button class="btn btn-primary mx-3" id="editBtn" type="button" id="button-addon2" onclick="editAssignment(this)">Edit</button>
    <button class="btn btn-danger " id="delBtn" type="button" id="button-addon2"  onclick="delAssignment(this)">Delete</button>`

    mainList.appendChild(newLi)
}
function delAssignment(currElement) {        t7
    currElement.parentElement.remove()
    if (mainList.children.length <= 0) {
        let newMsg = document.createElement('h4')

        newMsg.id = "emptyMsg"
        newMsg.classList.add("middle")
        newMsg.textContent = "You Can Chill Now. Nothing to do."

        mainList.appendChild(newMsg)
    }
}
function editAssignment(currElement) {
    if (currElement.textContent == 'Done') {
        currElement.textContent = 'Edit'
        
        let newtext = currElement.previousElementSibling.value
        
        let completedtext = document.createElement('h4')
        completedtext.className = 'flex-grow-1'
        completedtext.textContent = newtext 

        currElement.parentElement.replaceChild(completedtext, currElement.previousElementSibling)
    }
    else {
        currElement.textContent = 'Done'

        let editedText = document.createElement('input')
        editedText.type = 'text'
        editedText.className = 'form-control'
        editedText.placeholder = 'Assignment Name'
        editedText.maxLength = 40
        
        currElement.parentElement.replaceChild(editedText, currElement.previousElementSibling)
    }
}