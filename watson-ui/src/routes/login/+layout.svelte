<script lang="ts">
  import "../../global.css";

</script>

<main>
  <section>
    <div class="lside"> <span class="customw">Welcome back,</span></div>
    <div class="rside"> <div class="watson-logo"></div> <slot /> </div>
  </section>
</main>

<style>

  @import url('https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;500;600&family=Quicksand:wght@300;500&display=swap');
  main {
    font-family: 'Open Sans', sans-serif;
    margin: 0 auto;
    padding: 0;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    min-height: 100vh;
  }
  section {
    width: 50%;
    height: 400px;
    overflow: hidden;
    border-radius: 10px;
    display: flex;
    align-items: center;
    flex-direction: row;
    text-align: center;
    box-shadow: 2px 11px 8px -5px rgba(158,158,158,0.68);
    -webkit-box-shadow: 2px 11px 8px -5px rgba(158,158,158,0.68);
    -moz-box-shadow: 2px 11px 8px -5px rgba(158,158,158,0.68);
  }

  .lside {
    width: 50%;
    height: 100vh;
    background: radial-gradient(732px at 96.2% 89.9%, rgb(70, 66, 159) 0%, rgb(187, 43, 107) 92%);
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .rside {
    width: 50%;
    height: auto;
    background-color: rgb(255, 255, 255);
    padding: 5em;
    overflow: hidden;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
  }

  .watson-logo {
    width: 300px;
    height: 120px;
    display: block;
    background-repeat: no-repeat;
    background-size: contain;
    background-image: url('/watson-logo.png');
  }

  span.customw {
    font-size: 1.5em;
    background: #F6F8F9;
    background: radial-gradient(circle at 18.7% 37.8%, rgb(250, 250, 250) 0%, rgb(225, 234, 238) 90%);
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
  }
</style>