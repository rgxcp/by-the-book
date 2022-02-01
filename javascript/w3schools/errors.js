try {
  addNumber(5, 5);
} catch (err) {
  console.log(`Nama error: ${err.name} \nPesan error: ${err.message}`);
} finally {
  console.log("Finally di jalankan");
}
