<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Property Search</title>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

    @import url('https://fonts.googleapis.com/css2?family=Playfair+Display:wght@700&family=DM+Sans:wght@400;500&display=swap');

    body {
      font-family: 'DM Sans', sans-serif;
      background: #0d0d0d;
      color: #f0ede6;
      min-height: 100vh;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      padding: 2rem;
    }

    .hero {
      text-align: center;
      max-width: 600px;
      width: 100%;
    }

    .hero h1 {
      font-family: 'Playfair Display', serif;
      font-size: clamp(2.4rem, 6vw, 4rem);
      line-height: 1.1;
      margin-bottom: 1rem;
      color: #f0ede6;
      letter-spacing: -0.02em;
    }

    .hero h1 span {
      color: #c9a84c;
    }

    .hero p {
      color: #888;
      font-size: 1.05rem;
      margin-bottom: 2.5rem;
    }

    .search-form {
      display: flex;
      gap: 0.75rem;
      width: 100%;
    }

    .search-form input {
      flex: 1;
      padding: 0.9rem 1.2rem;
      font-size: 1rem;
      font-family: 'DM Sans', sans-serif;
      background: #1a1a1a;
      border: 1px solid #2e2e2e;
      border-radius: 8px;
      color: #f0ede6;
      outline: none;
      transition: border-color 0.2s;
    }

    .search-form input::placeholder { color: #555; }
    .search-form input:focus { border-color: #c9a84c; }

    .search-form button {
      padding: 0.9rem 1.8rem;
      background: #c9a84c;
      color: #0d0d0d;
      font-size: 1rem;
      font-weight: 600;
      font-family: 'DM Sans', sans-serif;
      border: none;
      border-radius: 8px;
      cursor: pointer;
      transition: background 0.2s;
      white-space: nowrap;
    }

    .search-form button:hover { background: #e0bb60; }

    footer {
      margin-top: 4rem;
      color: #333;
      font-size: 0.8rem;
    }
  </style>
</head>
<body>
  <div class="hero">
    <h1>Find <span>Properties</span><br/>Anywhere</h1>
    <p>Search a location to browse available rental properties.</p>

    {{/* Simple GET form — submits to /all/:location via query param redirect */}}
    <form class="search-form" action="/search" method="GET">
      <input
        type="text"
        name="location"
        placeholder="e.g. destin, usa, florida..."
        required
        autocomplete="off"
      />
      <button type="submit">Search</button>
    </form>
  </div>

  <footer>Powered by RentByOwner</footer>
</body>
</html>