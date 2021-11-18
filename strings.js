/**
 * Implement an algorithm to determine if a string has all unique characters.
 * What if you cannot use additional data structures?
 */
function isUnique (string = '') {
  const UNICODE_MAX = 144697
  const hash_map = []

  let i = 0
  for (i; i < UNICODE_MAX; i ++) {
    hash_map.push(0)
  }

  for (let i = 0; i < string.length; i ++) {
    const charUnicode = string[i].charCodeAt()
    const value = hash_map[charUnicode]
    if (value !== 1) hash_map[charUnicode] = 1
    else return false
  }
  
  return true
}

/**
 * Given two strings,write a method to decide if one is a permutation of the
 * other.
 */
function checkPermutation (stringA = '', stringB = '') {
  if (stringA.length !== stringB.length) return false

  const hashMap = []

  const UNICODE_MAX = 144697
  for (let i = 0; i < UNICODE_MAX; i ++) {
    hashMap.push(0)
  }

  for (let i = 0; i < stringA.length; i ++) {
    const unicodeA = stringA[i].charCodeAt()
    hashMap[unicodeA] += 1
  }

  for (let i = 0; i < stringB.length; i++) {
    const unicodeB = stringB[i].charCodeAt()
    hashMap[unicodeB] -= 1
    if (hashMap[unicodeB] < 0) return false
  }

  return true
}

/**
 * Write a method to replace all spaces in a string with '%20'.
 * You may assume that the string has sufficient space at the end to hold the additional characters,and that you are given the "true" length of the string.
 * (Note: If implementing in Java,please use a character array so that you can perform this operation in place.)
 */

function URLify (string, trueLength) {
  console.time('URLify')
  let charArr = []
  for (let i = 0; i < trueLength; i ++) {
    const unicode = string[i].charCodeAt()
    if (unicode === 32) charArr.push(`%20`)
    else charArr.push(string[i])
  }

  const result = charArr.join('')
  console.timeEnd('URLify')
  return result
}

// anna kayak
function palindromePermutation (string = '') {
  let charArray = []
  for (let i = 0; i < string.length; i ++) {
    if (string[i].charCodeAt() !== 32) charArray.push(string[i])
  }

  const UNICODE_MAX = 144697
  const hash_map = new Array(UNICODE_MAX).fill(0)

  const isEvenLength = charArray.length % 2 === 0

  for (let i = 0; i < charArray.length; i ++) {
    const charUnicode = charArray[i].charCodeAt()
    hash_map[charUnicode] += 1
  }

  if (isEvenLength) {
    for (let i = 0; i < charArray.length; i++) {
      const charUnicode = charArray[i].charCodeAt()
      if (hash_map[charUnicode] !== 2) return false
    }
  } else {
    const unEvenCount = []
    for (let i = 0; i < charArray.length; i ++) {
      const charUnicode = charArray[i].charCodeAt()
      if (hash_map[charUnicode] !== 2) unEvenCount.push(charArray[i])
    }

    return !(unEvenCount.length > 1)
  }

}

/**
 * There are three types of edits that can be performed on strings:
 * insert a character, remove a character, or replace a character.
 * Given two strings, write a function to check if they are one edit (or zero edits) away.
 */
function oneAway (string1 = '', string2 = '') {
  const hashMap = new Array(144697).fill(0)
  const diff = []

  for (let i = 0; i < string1.length; i ++) {
    const charUniCode = string1[i].charCodeAt()
    hashMap[charUniCode] += 1
  }

  for (let i = 0; i < string2.length; i ++) {
    const charUniCode = string2[i].charCodeAt()
    if (!(hashMap[charUniCode] > 0)) diff.push(string2[i])
  }

  return !(diff.length > 1)
}

/**
 * Implement a method to perform basic string compression using the counts of repeated characters.
 * For example, the string aabcccccaaa would become a2blc5a3.
 * If the "compressed" string would not become smaller than the original string, your method should return
 * the original string. You can assume the string has only uppercase and lowercase letters (a - z)
 */
function stringCompression (string = '') {
  let result = ''
  let charCount = 0

  for (let i = 0; i < string.length; i++) {
    const previos = string[i - 1]
    const current = string[i]
    const next = string[i + 1]

    if (previos !== current) charCount = 1
    if (previos === current) charCount++
    if (current !== next) result += `${current}${charCount}`
  }

  return result.length > string.length ? string : result
}

