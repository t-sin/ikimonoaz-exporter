package template

const indexTemplate = `
<html>
  <head>
    <title>{{ user.name }} - いきものAZ エクスポート結果</title>
    <style>
      /* global layout */
      * {
        margin: 0;
        padding: 0;
      }
      body {
        background-color: #f7f7f7;
        color: #333333;
        font-size: 1.1em;
      }
      header {
        padding: 30px;
      }
      footer {
        padding: 1em 0;
      }
      h1 {
        text-align: center;
      }
      #container {
        margin: 0 auto;
        max-width: 900px;
      }
      main {
        padding: 40px;
        background-color: #ffffff;
        box-shadow: 2px 2px 4px 4px rgba(0,0,0,0.1);
      }

      section {
        margin: 2em 0;
      }
      h2 {
        margin-bottom: 0.5em;
      }
      ul {
        margin-left: 30px;
        list-style: none;
      }
      h3 {
        margin-top: 1em;
      }
      dl {
        margin-left: 30px;
      }
      dt {
        font-weight: 500;
      }
      dd {
        margin-left: 1em;
        margin-bottom: 0.5em;
      }
      /*  layout */
      .meister::before {
        content: "🏅";
      }
      .tag {
        margin: 1px;
        padding: 1px 4px;
        border: solid 1px #afcdde;
        border-radius: 4px;
        background-color: #dff0f9;
        color: #1e4f6b;
        font-size: 0.8em;
      }
      .creatures {
        font-size: 0.9em;
      }
      .creature {
        padding: 1px 2px;
        color: #072b40;
      }
      .creature-place {
        font-size: 0.85em;
        color: #072b40;
      }
      .date {
        color: #777777;
        font-size: 0.85em;
      }
      .date span {
        font-size: 0.85em;
      }
    </style>
  </head>
  <body>
    <div id="container">

      <header>
        <h1>いきものAZ エクスポート結果</h1>
      </header>
    
      <main>
        <section id="user">
          <h2>ユーザ情報</h2>
          <dl>
            <dt>ユーザ名</dt><dd>{{ user.name }}</dd>
            <dt>プロフィール</dt><dd>{{ user.profile }}</dd>
            <dt>よく行く園館</dt><dd>{{ user.place }}</dd>
            <dt>マイスター</dt>
            <dd>
              <ul>
                {{ #user.meister }}
                <li class="meister">{{.}}</li>
                {{ /user.meister }}
              </ul>
            </dd>
          </dl>
        </section>

        <section id="article-list">
          <h2>投稿記事一覧</h2>
      
          <ul>
            {{ #articles }}
            <li>
              <article>
                <h3>{{ title }}</h3>
                <p class="creatures">
                  {{ #creatures }}
                  <span class="creature">{{ name }}</span><span class="creature-place">@{{ place }}</span>
                  {{ /creatures }}
                </p>
                <p class="tags">
                  {{ #tags }}
                  <span class="tag">{{ . }}</span>
                  {{ /tags }}
                </p>
                <p class="date">
                  <span>作成:</span>{{ created_at }},
                  <span>編集:</span>{{ updated_at }},
                  <span>公開:</span>{{ released_at }}
                </p>
              </article>
            </li>
            {{ /articles }}
          </ul>
        </section>
      </main>

      <footer>
        Generated by <a href="https://github.com/t-sin/ikimonoaz-exporter">ikimonoaz-exporter</a>
      </footer>

    </div>
  </body>
</html>
`
