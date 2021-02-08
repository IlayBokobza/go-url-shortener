//dom

//open popup
const newUrlBtn = document.getElementById('newUrlBtn')
const popup = document.getElementById('popup')

newUrlBtn.addEventListener('click',() => {
    popup.style.display = 'block'
})

//close popup
const closePopupBtn = document.getElementById('closePopupBtn')
closePopupBtn.addEventListener('click',e => {
    e.preventDefault()
    popup.style.display = 'none'
})




////////////////
////////////////
////////////////
////////////////
//http
const urlContainer = document.getElementById('urlContainer')

//gets data from server
fetch('/api/get').then(async res => {
    const data = await res.json()

    data.forEach(item => {
        urlContainer.innerHTML += `
        <div class="box">
            <p>URL: ${item.url}</p>
            <p>ID: ${item.id}</p>
            <span class="delete-icon" data-url="${item.url}">&#9850;</span>
        </div>`
    })

    //sets event listener for delete btn
    const deleteBtns = document.querySelectorAll('.delete-icon')
    deleteBtns.forEach((btn) => {
        btn.addEventListener('click',deleteUrl)
    })
})


//add new url
const newUrlForm = document.getElementById('newUrlForm')
const urlInput = document.getElementById('urlInput')
const newUrlFormFeedback = document.getElementById('newUrlFormFeedback')

newUrlForm.addEventListener('submit',async e => {
    e.preventDefault()
    const url = urlInput.value

    if(url.trim() === ""){
        newUrlFormFeedback.textContent = 'Fill Fields'
        return
    }

    const res = await fetch('/api/add',{
        method:'POST',
        mode:'cors',
        cache:'default',
        credentials:'same-origin',
        headers:{
            'content-type':'application/json',
        },
        redirect:"follow",
        referrerPolicy:'no-referrer',
        body:url
    }).catch(err => console.warn(err))

    try{
        let json = JSON.parse(await res.text())
        newUrlFormFeedback.textContent = ""

        urlContainer.innerHTML += `
        <div class="box">
            <p>URL: ${json.url}</p>
            <p>ID: ${json.id}</p>
            <span class="delete-icon" data-url="${json.url}">&#9850;</span>
        </div>`
    }catch{
        newUrlFormFeedback.textContent = "URL Already Stored"
        return
    }

})

//delete url
const deleteUrl = async (e) => {
    const url = e.target.dataset.url

    await fetch('/api/delete',{
        method:'DELETE',
        mode:'cors',
        cache:'default',
        credentials:'same-origin',
        redirect:"follow",
        referrerPolicy:'no-referrer',
        body:url
    }).catch(err => console.warn(err))

    location.reload()
}