// @ts-check

export function dayRate(ratePerHour) {
  return ratePerHour*8;
}

export function daysInBudget(budget, ratePerHour) {
  return Math.floor(budget/(ratePerHour*8));
}

export function priceWithMonthlyDiscount(ratePerHour, numDays, discount) {
  let remDays = numDays%22;
  let costRemDays=remDays*8*ratePerHour;
  let month=numDays-remDays;
  let monthCost=month*8*ratePerHour;
  let disc=monthCost*(1-discount);
  return Math.ceil(disc+costRemDays);
  
}
