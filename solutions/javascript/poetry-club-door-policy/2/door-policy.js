// @ts-check

export function frontDoorResponse(line) {
  return line[0];
}

export function frontDoorPassword(word) {
  return word[0].toUpperCase()+word.slice(1).toLowerCase();
}

export function backDoorResponse(line) {
  return line.trim().slice(-1);
}


export function backDoorPassword(word) {
  let first = word[0].toUpperCase();
  let rest = word.slice(1).toLowerCase();
  return `${frontDoorPassword(word)}, please`;
}
