const urlContainer = document.getElementById('urlContainer')

fetch('/api/get').then(async res => {
    const data = await res.json()

    data.forEach(item => {
        urlContainer.innerHTML += `
        <div class="box">
            <p>URL: ${item.url}</p>
            <p>ID: ${item.id}</p>
        </div>`
    })
})