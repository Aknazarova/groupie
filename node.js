fs = require('fs')
https = require('https')

getUrl('https://01.alem.school/api/object/', res => {
  console.log(res)

  let list = res.children['div-01'].children.exam.children

  let exs = []

  for (let i = 1; i <= 7; i++) {
    let level = getRandom(i, list)
    exs.push(level[Math.floor(Math.random() * level.length)])
  }

  console.log(exs)
})

const getRandom = (level, list) => {
  let ans = []

  for (let key in list) {
    if (list[key].attrs.group === level) {
      ans.push({
        name: key,
        link: list[key].attrs['subject-en'],
        level,
        allowed: list[key].attrs.allowedFunctions
      })
    }
  }

  return ans
}

function getUrl(url, callback) {
  https.get(url, res => {
    let rawData = ''

    res.on('data', chunk => {
      rawData += chunk
    })

    res.on('end', () => {
      try {
        const body = JSON.parse(rawData)
        callback(body)
      } catch (e) {
        console.error(e.message)
      }
    })
  })
}