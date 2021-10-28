<script lang="ts">
  import ChannelList from "../components/ChannelList.svelte";
  import MemberList from "../components/MemberList.svelte";
  import MessageBox from "../components/MessageBox.svelte";
</script>

<div class="chat">
  <section>
    <h3>Channels</h3>
    <hr />
    <ChannelList />
  </section>
  <section>
    <h3>Messages</h3>
    <hr />
    <MessageBox />
  </section>
  <section>
    <h3>Members</h3>
    <hr />
    <MemberList />
  </section>
</div>

<style>
  .chat {
    display: grid;
    grid-template-columns: 1fr 3fr 1fr;
    column-gap: 1em;
    padding: 0 3em;
  }
</style>
