import { m, useScroll, useSpring } from 'framer-motion';
// next
import Head from 'next/head';
// @mui
import { useTheme } from '@mui/material/styles';
import {Box, Container, Grid, Typography} from '@mui/material';
// sections
import {useSettingsContext} from "../components/settings";
import DashboardLayout from "../layouts/dashboard";
import { AppWidgetSummary} from "../sections/@dashboard/general/app";

// ----------------------------------------------------------------------

// HomePage.getLayout = (page: React.ReactElement) => <MainLayout> {page} </MainLayout>;
HomePage.getLayout = (page: React.ReactElement) => (
    <DashboardLayout>{page}</DashboardLayout>
);

// ----------------------------------------------------------------------

export default function HomePage() {
  const theme = useTheme();

  const { scrollYProgress } = useScroll();
  const { themeStretch } = useSettingsContext();

  const scaleX = useSpring(scrollYProgress, {
    stiffness: 100,
    damping: 30,
    restDelta: 0.001,
  });

  const progress = (
    <m.div
      style={{
        top: 0,
        left: 0,
        right: 0,
        height: 3,
        zIndex: 1999,
        position: 'fixed',
        transformOrigin: '0%',
        backgroundColor: theme.palette.primary.main,
        scaleX,
      }}
    />
  );

  return (
    <>
      <Head>
        <title>Estuary Metrics Explorer</title>
      </Head>

      {progress}

      {/*<HomeHero />*/}

      {/*<Box*/}
      {/*  sx={{*/}
      {/*    overflow: 'hidden',*/}
      {/*    position: 'relative',*/}
      {/*    bgcolor: 'background.default',*/}
      {/*  }}*/}
      {/*>*/}
          <Container maxWidth={themeStretch ? false : 'xl'}>
          <Typography variant="h4" sx={{ mb: 5 }}>
              Quick Stats
          </Typography>
          <Grid container spacing={3}>
              <Grid item xs={12} sm={6} md={3}>
                  <AppWidgetSummary
                      title="Total Root CIDs"
                      percent={2.6}
                      total={18765}
                      chart={{
                          colors: [theme.palette.primary.main],
                          series: [5, 18, 12, 51, 68, 11, 39, 37, 27, 20],
                      }}
                  />
              </Grid>
              <Grid item xs={12} sm={6} md={3}>
                  <AppWidgetSummary
                      title="Total Number of Object References"
                      percent={2.6}
                      total={18765}
                      chart={{
                          colors: [theme.palette.primary.main],
                          series: [5, 18, 12, 51, 68, 11, 39, 37, 27, 20],
                      }}
                  />
              </Grid>
              <Grid item xs={12} md={4}>
                  <AppWidgetSummary
                      title="Active Successful Storage Deals"
                      percent={2.6}
                      total={18765}
                      chart={{
                          colors: [theme.palette.primary.main],
                          series: [5, 18, 12, 51, 68, 11, 39, 37, 27, 20],
                      }}
                  />
              </Grid>
              <Grid item xs={12} md={4}>
                  <AppWidgetSummary
                      title="Total Pinned CIDs for Hot Retrieval"
                      percent={2.6}
                      total={18765}
                      chart={{
                          colors: [theme.palette.primary.main],
                          series: [5, 18, 12, 51, 68, 11, 39, 37, 27, 20],
                      }}
                  />
              </Grid>

              <Grid item xs={12} md={4}>
                  <AppWidgetSummary
                      title="Total Sealed Storage Contributed to Filecoin"
                      percent={2.6}
                      total={18765}
                      chart={{
                          colors: [theme.palette.primary.main],
                          series: [5, 18, 12, 51, 68, 11, 39, 37, 27, 20],
                      }}
                  />
              </Grid>

              <Grid item xs={12} md={4}>
                  <AppWidgetSummary
                      title="Total Storage Providers Receiving Deals from Estuary Node"
                      percent={2.6}
                      total={18765}
                      chart={{
                          colors: [theme.palette.primary.main],
                          series: [5, 18, 12, 51, 68, 11, 39, 37, 27, 20],
                      }}
                  />
              </Grid>
              <Grid item xs={12} md={4}>
                  <AppWidgetSummary
                      title="Total Registered Users"
                      percent={2.6}
                      total={18765}
                      chart={{
                          colors: [theme.palette.primary.main],
                          series: [5, 18, 12, 51, 68, 11, 39, 37, 27, 20],
                      }}
                  />
              </Grid>
          </Grid>
              <hr/>
              <br/>
              <Typography variant="h4" sx={{ mb: 5 }}>
                  Rates
              </Typography>
              <br/>
              <hr/>
              <Typography variant="h4" sx={{ mb: 5 }}>
                  Ranking Stats
              </Typography>
              <Grid item xs={12} md={6} lg={8}>
                  {/*<AppTopInstalledCountries title="Top Miners" list={_appInstalled} />*/}
              </Grid>
          </Container>
        {/*<HomeMinimal />*/}

        {/*<HomeHugePackElements />*/}

        {/*<HomeForDesigner />*/}

        {/*<HomeDarkMode />*/}

        {/*<HomeColorPresets />*/}

        {/*<HomeCleanInterfaces />*/}

        {/*<HomePricingPlans />*/}

        {/*<HomeLookingFor />*/}

        {/*<HomeAdvertisement />*/}
      {/*</Box>*/}
    </>
  );
}
