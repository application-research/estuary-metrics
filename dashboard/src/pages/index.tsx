import { m, useScroll, useSpring } from 'framer-motion';
// next
import Head from 'next/head';
// @mui
import { useTheme } from '@mui/material/styles';
import {Box, Container, Grid, Typography} from '@mui/material';
// layouts
import MainLayout from '../layouts/main';
// sections
import {
  HomeHero,
  HomeMinimal,
  HomeDarkMode,
  HomeLookingFor,
  HomeForDesigner,
  HomeColorPresets,
  HomePricingPlans,
  HomeAdvertisement,
  HomeCleanInterfaces,
  HomeHugePackElements,
} from '../sections/home';
import {AnalyticsWidgetSummary} from "../sections/@dashboard/general/analytics";
import {useSettingsContext} from "../components/settings";
import DashboardLayout from "../layouts/dashboard";
import GeneralAnalyticsPage from "./dashboard/analytics";

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
                  <AnalyticsWidgetSummary
                      title="TOTAL ROOT CIDS UPLOADED TO ESTUARY. THIS VALUE DOES NOT INCLUDE SUB OBJECTS REFERENCES."
                      total={714000}
                      icon={'ant-design:android-filled'}
                  />

              </Grid>
              <Grid item xs={12} sm={6} md={3}>
                  <AnalyticsWidgetSummary
                      title="TOTAL NUMBER OF OBJECT REFERENCES PROVIDED BY EVERY ROOT CID IN THE NETWORK."
                      total={714000}
                      icon={'ant-design:android-filled'}
                  />

              </Grid>
              <Grid item xs={12} sm={6} md={3}>
                  <AnalyticsWidgetSummary
                      title="ACTIVE SUCCESSFUL STORAGE DEALS ON THE FILECOIN NETWORK"
                      total={714000}
                      icon={'ant-design:android-filled'}
                  />

              </Grid>
              <Grid item xs={12} sm={6} md={3}>
                  <AnalyticsWidgetSummary
                      title="TOTAL PINNED IPFS STORAGE FOR HOT RETRIEVAL FROM ANY IPFS GATEWAY. THIS DATA IS NOT STORED ON FILECOIN"
                      total={714000}
                      icon={'ant-design:android-filled'}
                  />

              </Grid>
              <Grid item xs={12} sm={6} md={3}>
                  <AnalyticsWidgetSummary
                      title="TOTAL SEALED STORAGE CONTRIBUTED TO FILECOIN INCLUDING A 6X REPLICATION"
                      total={714000}
                      icon={'ant-design:android-filled'}
                  />

              </Grid>
              <Grid item xs={12} sm={6} md={3}>
                  <AnalyticsWidgetSummary
                      title="TOTAL STORAGE PROVIDERS RECEIVING DEALS FROM OUR ESTUARY NODE"
                      total={714000}
                      icon={'ant-design:android-filled'}
                  />

              </Grid>
              <Grid item xs={12} sm={6} md={3}>
                  <AnalyticsWidgetSummary
                      title="TOTAL REGISTERED USERS"
                      total={714000}
                      icon={'ant-design:android-filled'}
                  />

              </Grid>

          </Grid>
          <br/>
              <Typography variant="h4" sx={{ mb: 5 }}>
                  Quick Stats
              </Typography>
              <Grid container spacing={3}>
                  <Grid item xs={12} sm={6} md={3}>
                      <AnalyticsWidgetSummary
                          title="TOTAL ROOT CIDS UPLOADED TO ESTUARY. THIS VALUE DOES NOT INCLUDE SUB OBJECTS REFERENCES."
                          total={714000}
                          icon={'ant-design:android-filled'}
                      />

                  </Grid>
                  <Grid item xs={12} sm={6} md={3}>
                      <AnalyticsWidgetSummary
                          title="TOTAL NUMBER OF OBJECT REFERENCES PROVIDED BY EVERY ROOT CID IN THE NETWORK."
                          total={714000}
                          icon={'ant-design:android-filled'}
                      />

                  </Grid>
                  <Grid item xs={12} sm={6} md={3}>
                      <AnalyticsWidgetSummary
                          title="ACTIVE SUCCESSFUL STORAGE DEALS ON THE FILECOIN NETWORK"
                          total={714000}
                          icon={'ant-design:android-filled'}
                      />

                  </Grid>
                  <Grid item xs={12} sm={6} md={3}>
                      <AnalyticsWidgetSummary
                          title="TOTAL PINNED IPFS STORAGE FOR HOT RETRIEVAL FROM ANY IPFS GATEWAY. THIS DATA IS NOT STORED ON FILECOIN"
                          total={714000}
                          icon={'ant-design:android-filled'}
                      />

                  </Grid>
                  <Grid item xs={12} sm={6} md={3}>
                      <AnalyticsWidgetSummary
                          title="TOTAL SEALED STORAGE CONTRIBUTED TO FILECOIN INCLUDING A 6X REPLICATION"
                          total={714000}
                          icon={'ant-design:android-filled'}
                      />

                  </Grid>
                  <Grid item xs={12} sm={6} md={3}>
                      <AnalyticsWidgetSummary
                          title="TOTAL STORAGE PROVIDERS RECEIVING DEALS FROM OUR ESTUARY NODE"
                          total={714000}
                          icon={'ant-design:android-filled'}
                      />

                  </Grid>
                  <Grid item xs={12} sm={6} md={3}>
                      <AnalyticsWidgetSummary
                          title="TOTAL REGISTERED USERS"
                          total={714000}
                          icon={'ant-design:android-filled'}
                      />

                  </Grid>

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
