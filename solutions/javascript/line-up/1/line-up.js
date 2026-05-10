// @ts-check
// This is only a SKELETON file for the 'Line Up' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const format = (name,number) => {
  let ind = 'aa';
  if (number%100>= 11 && number%100 <=13){
    ind ='th';
  }else if (number%10===1){
    ind ='st';
  }else if (number%10===2){
    ind='nd';
  }else if (number%10===3){
    ind='rd';
  }else{
    ind='th';
  }
  return `${name}, you are the ${number}${ind} customer we serve today. Thank you!`
}
