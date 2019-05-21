package words.web.auth;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.runners.MockitoJUnitRunner;
import org.springframework.mock.web.MockHttpSession;

import javax.servlet.FilterChain;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;
import static org.springframework.http.HttpStatus.UNAUTHORIZED;

@RunWith(MockitoJUnitRunner.class)
public class AuthFilterTest {

    @Mock AuthService authService;
    @Mock HttpServletRequest request;
    @Mock HttpServletResponse response;
    @Mock FilterChain filterChain;
    MockHttpSession session = new MockHttpSession();

    @InjectMocks AuthFilter authFilter;

    @Before
    public void setUp() throws Exception {
        when(request.getSession(true)).thenReturn(session);
    }

    @Test
    public void fail_if_no_header() throws Exception {

        authFilter.doFilter(request, response, filterChain);

        verify(response).sendError(UNAUTHORIZED.value(), UNAUTHORIZED.getReasonPhrase());
    }

    @Test
    public void ok_if_there_is_header_and_user() throws Exception {
        when(request.getHeader("Authorization")).thenReturn("Basic YWRtaW46YWRtaW4=");
        when(authService.isUserExist("admin", "admin")).thenReturn(true);

        authFilter.doFilter(request, response, filterChain);
        verify(filterChain).doFilter(request, response);
    }

    @Test
    public void fail_if_no_user_exist() throws Exception {
        when(request.getHeader("Authorization")).thenReturn("Basic YWRtaW46YWRtaW4=");
        when(authService.isUserExist("admin", "admin")).thenReturn(false);

        authFilter.doFilter(request, response, filterChain);
        verify(response).sendError(UNAUTHORIZED.value(), UNAUTHORIZED.getReasonPhrase());
    }

    @Test
    public void wrong_content_of_filter() throws Exception {
        when(request.getHeader("Authorization")).thenReturn("blablabla");

        authFilter.doFilter(request, response, filterChain);
        verify(response).sendError(UNAUTHORIZED.value(), UNAUTHORIZED.getReasonPhrase());
    }


    @Test
    public void set_auth_info_if_there_is_header_and_user() throws Exception {
        when(request.getHeader("Authorization")).thenReturn("Basic YWRtaW46YWRtaW4=");
        when(authService.isUserExist("admin", "admin")).thenReturn(true);

        authFilter.doFilter(request, response, filterChain);
        assertThat(session.getAttribute("username"), is("admin"));
    }
}
