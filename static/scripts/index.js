let progressBar1 = document.querySelector(".circular-progress-1");
let valueContainer1 = document.querySelector(".value-container-1");

let progressValue1 = 0;
let progressEndValue1 = 32;
let speed1 = 50;

let progress = setInterval(() => {
  progressValue1++;
  valueContainer1.textContent = `${progressValue1}%`;
  progressBar1.style.background = `conic-gradient(
      #4d5bf9 ${progressValue1 * 3.6}deg,
      #cadcff ${progressValue1 * 3.6}deg
  )`;
  if (progressValue1 == progressEndValue1) {
    clearInterval(progress);
  }
}, speed1);


let progressBar2 = document.querySelector(".circular-progress-2");
let valueContainer2 = document.querySelector(".value-container-2");

let progressValue2 = 0;
let progressEndValue2 = 37;
let speed2 = 50;

let progress2= setInterval(() => {
  progressValue2++;
  valueContainer2.textContent = `${progressValue2}%`;
  progressBar2.style.background = `conic-gradient(
      #4d5bf9 ${progressValue2 * 3.6}deg,
      #cadcff ${progressValue2 * 3.6}deg
  )`;
  if (progressValue2 == progressEndValue2) {
    clearInterval(progress2);
  }
}, speed2);

// let progressBar3 = document.querySelector(".circular-progress-3");
// let valueContainer3 = document.querySelector(".value-container-3");

// let progressValue3 = 0;
// let progressEndValue3 = 32;
// let speed3 = 50;

// let progress3= setInterval(() => {
//   progressValue3++;
//   valueContainer3.textContent = `${progressValue3}%`;
//   progressBar3.style.background = `conic-gradient(
//       #4d5bf9 ${progressValue3* 3.6}deg,
//       #cadcff ${progressValue3* 3.6}deg
//   )`;
//   if (progressValue3 == progressEndValue3) {
//     clearInterval(progress3);
//   }
// }, speed3);





