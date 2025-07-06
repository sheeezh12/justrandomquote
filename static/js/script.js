function getQuote() {
  fetch('/quote', {
    headers: { 'X-Quote-Key': 'iowehniuvht4iuhv8t3489tv82pqniwqvt8q4yvtiqyn8tn4yt83nt' }
  })
    .then(response => response.json())
    .then(data => {
      document.getElementById('quote-text').textContent = `"${data.text}"`;
      document.getElementById('quote-author').textContent = `— ${data.author}`;
    })
    .catch(error => {
      console.error('Gagal ambil kutipan:', error);
      document.getElementById('quote-text').textContent = 'Ups! Tidak bisa ambil kutipan.';
      document.getElementById('quote-author').textContent = '';
    });
}


window.onload = getQuote;