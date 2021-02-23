const useIsEnabled = (_: string) => true;

// NotEnabled Component
const NotEnabled = () => {
  const flag1 = useIsEnabled("someFlag");
  const flag2 = useIsEnabled("anotherFlag");
  return flag1 ? <h1>flag1 enabled</h1> : flag2 ? <h1>flag2 enabled</h1> : <h1>all features disabled</h1>
}

const App = () => {
  const flag = useIsEnabled("someFlag");
  return <div>
    {flag ? <h1>Hello</h1> : <NotEnabled/>}
  </div>
}

App();