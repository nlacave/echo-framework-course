const getData = async() => {
    const req = await fetch("http://localhost:8888/data")
    const data = await req.json()
    drawData(data)
}

const drawData = data => {
    data.forEach(e => {
        const html = `<div>${e.name}</div>`
        content.insertAdjacentHTML('beforeend', html)
    });
}

obtener.addEventListener('click', e => {
    e.preventDefault()
    getData()
})