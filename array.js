function flattenArray (arr = []) {
  const accumulator = []
  flatten(arr, accumulator)
  return accumulator
}

function flatten (arr = [], accumulator = []) {
  arr.forEach(value => {
    if (Array.isArray(value)) flatten(value, accumulator)
    else accumulator.push(value)
  })
}
const result = flattenArray([0,[1,2], 1,2,3,[4,5,[6], [7,8,9,[10, 11]]]])
console.log(result)




