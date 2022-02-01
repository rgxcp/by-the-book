let day;
switch (new Date().getDay()) {
  case 0:
    day = "Minggu";
    break; // berhenti
  case 1:
    day = "Senin";
    break; // berhenti
  case 2:
    day = "Selasa";
    break; // berhenti
  case 3:
    day = "Rabu";
    break; // berhenti
  case 4:
    day = "Kamis";
    break; // berhenti
  case 5:
    day = "Jumat";
    break; // berhenti
  case 6:
    day = "Sabtu";
}
console.log(day);

const score = 80;
let message;
switch (score) {
  case 50:
    message = "D";
    break;
  case 60:
  case 70:
    message = "C";
    break;
  case 80:
  case 90:
    message = "B";
    break;
  case 100:
    message = "A";
    break;
  default:
    message = "E";
}
console.log(message);
