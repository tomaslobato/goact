const esbuild = require("esbuild");

async function build() {
  const ctx = await esbuild.context({
    entryPoints: ["client/main.tsx", "client/styles.css"],
    outdir: "dist",
    bundle: true,
    minify: true,
    logLevel: "info",
  });

  await ctx.watch();
  console.log("⚡ Build complete! ⚡");
}

build().catch((error) => {
  console.error("Build failed:", error);
  process.exit(1);
});
