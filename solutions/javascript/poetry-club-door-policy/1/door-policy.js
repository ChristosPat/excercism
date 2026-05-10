// @ts-check

export function frontDoorResponse(line) {
  return line[0];
}

export function frontDoorPassword(word) {
  let first = word[0].toUpperCase();
  let rest = word.slice(1).toLowerCase();
  return first+rest;
}

export function backDoorResponse(line) {
  let pass=line.trim();
  let count = pass.slice(-1);
  return count;
}


export function backDoorPassword(word) {
  let first = word[0].toUpperCase();
  let rest = word.slice(1).toLowerCase();
  return first+rest+', please';
}
