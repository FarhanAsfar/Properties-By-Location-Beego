<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>Properties in {{.Location}}</title>
  <link rel="preconnect" href="https://fonts.googleapis.com"/>
  <link href="https://fonts.googleapis.com/css2?family=Playfair+Display:wght@700&family=DM+Sans:wght@400;500&display=swap" rel="stylesheet"/>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

    body {
      font-family: 'DM Sans', sans-serif;
      background: #0d0d0d;
      color: #f0ede6;
      min-height: 100vh;
      padding: 3rem 2rem;
    }

    .container { max-width: 1100px; margin: 0 auto; }

    /* ── Header ── */
    .page-header {
      display: flex;
      align-items: baseline;
      justify-content: space-between;
      flex-wrap: wrap;
      gap: 1rem;
      margin-bottom: 2.5rem;
      padding-bottom: 1.5rem;
      border-bottom: 1px solid #1e1e1e;
    }

    .page-header h1 {
      font-family: 'Playfair Display', serif;
      font-size: clamp(1.8rem, 4vw, 2.8rem);
      letter-spacing: -0.02em;
    }

    .page-header h1 span { color: #c9a84c; }

    .back-link { color: #888; text-decoration: none; font-size: 0.9rem; }
    .back-link:hover { color: #c9a84c; }

    /* ── Error state ── */
    .error-box {
      background: #1a0f0f;
      border: 1px solid #5c2222;
      border-radius: 10px;
      padding: 2rem;
      color: #e07070;
      text-align: center;
    }
    .error-box a { display: inline-block; margin-top: 1rem; color: #c9a84c; text-decoration: none; }

    /* ── Empty state ── */
    .empty-box { text-align: center; padding: 4rem 2rem; color: #555; }
    .empty-box p { font-size: 1.1rem; margin-bottom: 1rem; }

    /* ── Cards grid ── */
    .cards-grid {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
      gap: 1.5rem;
    }

    .card {
      background: #141414;
      border: 1px solid #1e1e1e;
      border-radius: 12px;
      overflow: hidden;
      display: flex;
      flex-direction: column;
      transition: border-color 0.2s, transform 0.2s;
    }

    .card:hover { border-color: #c9a84c44; transform: translateY(-2px); }

    /* Card top row: type badge + price */
    .card-header {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 1.2rem 1.2rem 0;
    }

    .card-badge {
      font-size: 0.7rem;
      font-weight: 600;
      letter-spacing: 0.08em;
      text-transform: uppercase;
      color: #c9a84c;
      background: #c9a84c18;
      border-radius: 4px;
      padding: 0.25rem 0.6rem;
    }

    .card-price { font-size: 1rem; font-weight: 600; color: #f0ede6; }
    .card-price span { font-size: 0.75rem; font-weight: 400; color: #555; }

    /* Card body */
    .card-body { padding: 1rem 1.2rem; flex: 1; }

    .card-body h2 {
      font-family: 'Playfair Display', serif;
      font-size: 1.05rem;
      line-height: 1.35;
      margin-bottom: 0.5rem;
      color: #f0ede6;
    }

    .card-location { font-size: 0.82rem; color: #666; margin-bottom: 0.9rem; }

    /* Stats row: beds / baths / guests */
    .card-stats {
      display: flex;
      gap: 1rem;
      font-size: 0.82rem;
      color: #888;
      margin-bottom: 0.9rem;
      flex-wrap: wrap;
    }

    /* Review row */
    .card-review { display: flex; align-items: center; gap: 0.5rem; font-size: 0.82rem; color: #888; }

    .review-score {
      background: #1e3a1e;
      color: #6fcf6f;
      font-size: 0.78rem;
      font-weight: 600;
      padding: 0.2rem 0.5rem;
      border-radius: 4px;
    }

    /* Card footer: booking CTA */
    .card-footer { padding: 0.9rem 1.2rem; border-top: 1px solid #1a1a1a; }

    .card-footer a {
      display: block;
      text-align: center;
      padding: 0.6rem;
      background: #c9a84c;
      color: #0d0d0d;
      font-size: 0.85rem;
      font-weight: 600;
      border-radius: 6px;
      text-decoration: none;
      transition: background 0.2s;
    }
    .card-footer a:hover { background: #e0bb60; }

    footer { margin-top: 4rem; text-align: center; color: #2a2a2a; font-size: 0.8rem; }
  </style>
</head>
<body>
  <div class="container">

    <div class="page-header">
      <h1>Properties in <span>{{.Location}}</span></h1>
      <a class="back-link" href="/">&#8592; New search</a>
    </div>

    {{if .Error}}
      {{/* Server-rendered error — no JS needed */}}
      <div class="error-box">
        <p>{{.Error}}</p>
        <a href="/">&#8592; Try a different location</a>
      </div>

    {{else if not .HasProperties}}
      <div class="empty-box">
        <p>No properties found for <strong>{{.Location}}</strong>.</p>
        <a class="back-link" href="/">Try a different location</a>
      </div>

    {{else}}
      <div class="cards-grid">
        {{range .Properties}}
        <div class="card">

          <div class="card-header">
            {{if .PropertyType}}
              <span class="card-badge">{{.PropertyType}}</span>
            {{else}}
              <span class="card-badge">Property</span>
            {{end}}
            {{if .Price}}
              <span class="card-price">${{printf "%.0f" .Price}} <span>/ night</span></span>
            {{end}}
          </div>

          <div class="card-body">
            <h2>{{if .PropertyName}}{{.PropertyName}}{{else}}Unnamed Property{{end}}</h2>

            {{if .Display}}
              <p class="card-location">&#128205; {{.Display}}</p>
            {{else if .City}}
              <p class="card-location">&#128205; {{.City}}{{if .Country}}, {{.Country}}{{end}}</p>
            {{end}}

            <div class="card-stats">
              {{if .Bedrooms}}<span>&#127beds; {{.Bedrooms}} bed{{if gt .Bedrooms 1}}s{{end}}</span>{{end}}
              {{if .Bathrooms}}<span>&#x1F6C1; {{.Bathrooms}} bath{{if gt .Bathrooms 1}}s{{end}}</span>{{end}}
              {{if .Occupancy}}<span>&#128101; Up to {{.Occupancy}} guests</span>{{end}}
            </div>

            {{if .ReviewScore}}
              <div class="card-review">
                <span class="review-score">{{printf "%.0f" .ReviewScore}}</span>
                <span>Review score</span>
                {{if .StarRating}}&nbsp;&#183;&nbsp;{{.StarRating}}&#9733;{{end}}
              </div>
            {{end}}
          </div>

          {{if .URL}}
          <div class="card-footer">
            <a href="{{.URL}}" target="_blank" rel="noopener noreferrer">View Property &rarr;</a>
          </div>
          {{end}}

        </div>
        {{end}}
      </div>
    {{end}}

  </div>

  <footer>Powered by RentByOwner</footer>
</body>
</html>