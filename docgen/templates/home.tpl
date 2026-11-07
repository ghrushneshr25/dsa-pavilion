---
slug: /
title: Home
hide_title: true
displayed_sidebar: null
description: Companion notes and Go implementations for Data Structures and Algorithms Made Easy by Narasimha Karumanchi.
---

<div className="dsa-home dsa-home--book">
  <header className="dsa-hero">
    <div className="dsa-hero__rail">
      <span>ctrl.plane // dsa</span>
      <span>Go implementations</span>
    </div>
    <div className="dsa-hero__layout">
      <div>
        <div className="dsa-kicker">Catalog online</div>
        <h1 className="dsa-hero__title">
          Data Structures
          <span>And Algorithms</span>
        </h1>
        <div className="dsa-hero__role">Made Easy · Narasimha Karumanchi</div>
        <div className="dsa-hero__lede">
          Companion notes and Go implementations. Open a topic to browse concepts and problems side by side.
        </div>
        <div className="dsa-hero__actions">
          <a className="dsa-btn dsa-btn--solid" href="#topics-heading">Browse topics</a>
          <a className="dsa-btn" href="https://github.com/ghrushneshr25/dsa" target="_blank" rel="noreferrer">GitHub</a>
        </div>
      </div>
      <aside className="dsa-status" aria-label="Catalog status">
        <div className="dsa-status__head">
          <span>sys.status</span>
          <span>godsa</span>
        </div>
        <div className="dsa-status__row">
          <span className="dsa-dot" aria-hidden="true"></span>
          <span className="dsa-status__label">topics</span>
          <span className="dsa-status__state">{{.Topics}}</span>
        </div>
        <div className="dsa-status__row">
          <span className="dsa-dot" aria-hidden="true"></span>
          <span className="dsa-status__label">concepts</span>
          <span className="dsa-status__state">{{.Concepts}}</span>
        </div>
        <div className="dsa-status__row">
          <span className="dsa-dot" aria-hidden="true"></span>
          <span className="dsa-status__label">problems</span>
          <span className="dsa-status__state">{{.Problems}}</span>
        </div>
        <div className="dsa-status__row">
          <span className="dsa-dot" aria-hidden="true"></span>
          <span className="dsa-status__label">runtime</span>
          <span className="dsa-status__state">GO</span>
        </div>
      </aside>
    </div>
  </header>
  <section className="dsa-catalog" aria-labelledby="topics-heading">
    <header className="dsa-section-header">
      <div className="dsa-section-meta">
        <span className="dsa-section-meta__index">01</span>
        <span>catalog</span>
        <span className="dsa-section-meta__rule" aria-hidden="true"></span>
      </div>
      <h2 id="topics-heading" className="dsa-section-title">Topics I am working through.</h2>
    </header>
    <div className="dsa-home__grid">
{{- range $i, $cat := .Categories }}
      <a className="dsa-home__card" href="./{{$cat.Slug}}">
        <span className="dsa-home__card-num">{{printf "%02d" (inc $i)}}</span>
        <div className="dsa-home__card-body">
          <div className="dsa-home__card-top">
            <span className="dsa-home__card-title">{{$cat.Name}}</span>
            <span className="dsa-home__card-arrow" aria-hidden="true">→</span>
          </div>
          <div className="dsa-home__card-stats">
            <span className="dsa-home__pill">{{ len $cat.Concepts }} concept{{ if ne (len $cat.Concepts) 1 }}s{{ end }}</span>
            <span className="dsa-home__pill dsa-home__pill--accent">{{ len $cat.Problems }} problem{{ if ne (len $cat.Problems) 1 }}s{{ end }}</span>
          </div>
        </div>
      </a>
{{- end }}
    </div>
  </section>
</div>
